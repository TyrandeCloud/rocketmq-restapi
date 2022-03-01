package rocketmq_restapi

type Message struct {
	MessageTag  string
	Properties  map[string]string
	MessageBody string
	MessageKey  string
}

type SendResponse struct {
	MessageID string
}

type WatchMessage struct {
	MessageTag       string
	Properties       map[string]string
	MessageID        string
	PublishTime      string
	ConsumedTimes    string
	FirstConsumeTime string
	NextConsumeTime  string
	MessageBody      string
	MessageKey       string
	// 消息句柄
	ReceiptHandle string
}
