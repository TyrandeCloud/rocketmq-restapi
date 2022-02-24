package rocketmq_restapi

import (
	"fmt"
	"testing"

	mq_http_sdk "github.com/aliyunmq/mq-http-go-sdk"
	"github.com/stretchr/testify/assert"
)

func mockClient(t *testing.T) *Client {
	opt, err := GetOptionFromEnv()
	assert.NoError(t, err)
	if err != nil {
		panic(err)
	}
	return NewClient(opt)
}

func TestClient(t *testing.T) {
	client := mockClient(t)

	tag := "TestClient"

	var msg mq_http_sdk.PublishMessageRequest
	msg = mq_http_sdk.PublishMessageRequest{
		MessageBody: "hello mq!",
		MessageTag:  tag,
		Properties: map[string]string{
			"pid": "123",
		},
	}
	ret, err := client.Producer().PublishMessage(msg)
	assert.NoError(t, err)
	if err == nil {
		fmt.Printf("Publish ---->\n\tMessageId:%s, BodyMD5:%s, \n", ret.MessageId, ret.MessageBodyMD5)
	}

	endChan := make(chan int)
	respChan := make(chan mq_http_sdk.ConsumeMessageResponse)
	errChan := make(chan error)
	go func() {
		client.Consumer(tag).ConsumeMessage(respChan, errChan, 3, 3)
		<-endChan
	}()
	resp := <-respChan
	//
	var handles []string
	fmt.Printf("Consume %d messages---->\n", len(resp.Messages))
	for _, v := range resp.Messages {
		handles = append(handles, v.ReceiptHandle)
		fmt.Printf("\tMessageID: %s, PublishTime: %d, MessageTag: %s\n"+
			"\tConsumedTimes: %d, FirstConsumeTime: %d, NextConsumeTime: %d\n"+
			"\tBody: %s\n"+
			"\tProps: %s\n",
			v.MessageId, v.PublishTime, v.MessageTag, v.ConsumedTimes,
			v.FirstConsumeTime, v.NextConsumeTime, v.MessageBody, v.Properties)
	}
	// ack
	err = client.Consumer(tag).AckMessage(handles)
	assert.NoError(t, err)
}
