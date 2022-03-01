package rocketmq_restapi

import (
	"context"
	"fmt"
	"strings"
	"time"

	mq_http_sdk "github.com/aliyunmq/mq-http-go-sdk"
	"github.com/gogap/errors"
)

type WatcherOption struct {
	Client        *Client
	Namespace     string
	Tag           string
	NumOfMessages int32
	WaitSeconds   int64
}

type Watcher struct {
	ctx context.Context
	cancel context.CancelFunc
	opt *WatcherOption
	ch  chan *WatchMessage
}

func NewWatcher(ctx context.Context, opt *WatcherOption) *Watcher {
	w := &Watcher{
		ctx: ctx,
		opt: opt,
	}
	return w
}

func (w *Watcher) watch() {
	if w.opt.NumOfMessages == 0 {
		w.opt.NumOfMessages = mq_http_sdk.DefaultNumOfMessages
	}
	errChan := make(chan error, 100)
	msgChan := make(chan mq_http_sdk.ConsumeMessageResponse, 4096)
	for {
		go w.opt.Client.Consumer(w.opt.Tag).ConsumeMessage(msgChan, errChan, w.opt.NumOfMessages, w.opt.WaitSeconds)
		select {
		case <-w.ctx.Done():
			return
		case err := <-errChan:
			if strings.Contains(err.(errors.ErrCode).Error(), "MessageNotExist") {
				fmt.Println("No new message, continue!")
			} else {
				fmt.Println(err)
				time.Sleep(time.Duration(3) * time.Second)
			}
		case msg := <-msgChan:
			w.handle(msg)
		case <-time.After(time.Second * time.Duration(w.opt.Client.Option.Timeout)):
		}
	}
}

func (w *Watcher) handle(msg mq_http_sdk.ConsumeMessageResponse) {
	for _, v := range msg.Messages {
		_ =v
	}
}

func (w *Watcher) Close() {
	w.cancel()
}
