package rocketmq_restapi

import (
	"fmt"
	"reflect"

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

func inspect(f interface{}) map[string]string {
	m := make(map[string]string)
	val := reflect.ValueOf(f).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		f := valueField.Interface()
		val := reflect.ValueOf(f)
		m[typeField.Name] = val.String()
	}
	return m
}

func (opt *Option) Validate() error {
	values := inspect(*opt)
	for k, v := range values {
		if v == "" {
			return fmt.Errorf("error: rocketmq option field %s was empty", k)
		}
	}
	return nil
}

type Client struct {
	mqc mq_http_sdk.MQClient
	opt *Option
}

func NewClient(opt *Option) *Client {
	c := &Client{
		mqc: mq_http_sdk.NewAliyunMQClient(opt.Endpoint, opt.AccessKeyId, opt.AccessKeySecret, opt.SecurityToken),
		opt: opt,
	}
	return c
}

func (c *Client) Producer() mq_http_sdk.MQProducer {
	return c.mqc.GetProducer(c.opt.InstanceId, c.opt.TopicName)
}

func (c *Client) Consumer(tag string) mq_http_sdk.MQConsumer {
	return c.mqc.GetConsumer(c.opt.InstanceId, c.opt.TopicName, c.opt.GroupId, tag)
}
