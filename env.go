package rocketmq_restapi

import (
	"errors"
	"fmt"
	"os"
)

const (
	AccessEndpoint  = "ALIYUN_ROCKETMQ_ACCESS_ENDPOINT"
	AccessKeyID     = "ALIYUN_ROCKETMQ_ACCESS_KEY_ID"
	AccessKeySecret = "ALIYUN_ROCKETMQ_ACCESS_SECRET"
	InstanceId      = "ALIYUN_ROCKETMQ_INSTANCE_ID"
	TopicName       = "ALIYUN_ROCKETMQ_TOPIC_NAME"
	GroupId         = "ALIYUN_ROCKETMQ_GROUP_ID"
	SecurityToken   = "ALIYUN_ROCKETMQ_SECURITY_TOKEN"
)

var errorTemplate = "unable to load configuration, %s must be defined"

var (
	ErrNoAccessEndpoint  = errors.New(fmt.Sprintf(errorTemplate, AccessEndpoint))
	ErrNoAccessKeyID     = errors.New(fmt.Sprintf(errorTemplate, AccessKeyID))
	ErrNoAccessKeySecret = errors.New(fmt.Sprintf(errorTemplate, AccessKeySecret))
	ErrNoInstanceId      = errors.New(fmt.Sprintf(errorTemplate, InstanceId))
	ErrNoTopicId         = errors.New(fmt.Sprintf(errorTemplate, TopicName))
	ErrNoGroupId         = errors.New(fmt.Sprintf(errorTemplate, GroupId))
	ErrNoSecurityToken   = errors.New(fmt.Sprintf(errorTemplate, SecurityToken))
)

func GetOptionFromEnv() (*Option, error) {
	opt := &Option{
		Endpoint:        os.Getenv(AccessEndpoint),
		AccessKeyId:     os.Getenv(AccessKeyID),
		AccessKeySecret: os.Getenv(AccessKeySecret),
		SecurityToken:   os.Getenv(SecurityToken),
		InstanceId:      os.Getenv(InstanceId),
		TopicName:       os.Getenv(TopicName),
		GroupId:         os.Getenv(GroupId),
	}
	if opt.Endpoint == "" {
		return nil, ErrNoAccessEndpoint
	}
	if opt.AccessKeyId == "" {
		return nil, ErrNoAccessKeyID
	}
	if opt.AccessKeySecret == "" {
		return nil, ErrNoAccessKeySecret
	}
	if opt.SecurityToken == "" {
		return nil, ErrNoSecurityToken
	}
	if opt.InstanceId == "" {
		return nil, ErrNoInstanceId
	}
	if opt.TopicName == "" {
		return nil, ErrNoTopicId
	}
	if opt.GroupId == "" {
		return nil, ErrNoGroupId
	}
	return opt, nil
}
