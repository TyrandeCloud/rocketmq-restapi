package rocketmq_restapi

import (
	"fmt"
	"github.com/aliyunmq/mq-http-go-sdk"
)

type Option struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string

	InstanceId string
	TopicName  string
	GroupId    string
}

type Client struct {
	mqc       mq_http_sdk.MQClient
	opt       *Option
	namespace string
}

func NewClient(opt *Option, namespace string) *Client {
	if namespace == "" {
		panic(fmt.Sprintf("namespace must be specified"))
	}
	c := &Client{
		mqc:       mq_http_sdk.NewAliyunMQClient(opt.Endpoint, opt.AccessKeyId, opt.AccessKeySecret, opt.SecurityToken),
		opt:       opt,
		namespace: namespace,
	}
	return c
}

func (c *Client) Tag(tag string) string {
	return fmt.Sprintf("%s:%s", c.namespace, tag)
}

func (c *Client) Producer() mq_http_sdk.MQProducer {
	return c.mqc.GetProducer(c.opt.InstanceId, c.opt.TopicName)
}

func (c *Client) Publish(tag string, msg mq_http_sdk.PublishMessageRequest) (resp mq_http_sdk.PublishMessageResponse, err error) {
	msg.MessageTag = c.Tag(tag)
	return c.mqc.GetProducer(c.opt.InstanceId, c.opt.TopicName).PublishMessage(msg)
}

func (c *Client) Consumer(tag string) mq_http_sdk.MQConsumer {
	return c.mqc.GetConsumer(c.opt.InstanceId, c.opt.TopicName, c.opt.GroupId, c.Tag(tag))
}
