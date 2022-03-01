package main

import (
	"flag"
	"fmt"
	"github.com/Shanghai-Lunara/pkg/zaplogger"
	rocketmqrestapi "github.com/TyrandeCloud/rocketmq-restapi"
	"github.com/TyrandeCloud/signals/pkg/signals"
	mq_http_sdk "github.com/aliyunmq/mq-http-go-sdk"
	"strconv"
	"time"
)

func main() {
	flag.Parse()
	stopCh := signals.SetupSignalHandler()
	zaplogger.Sugar().Info("game/server is running")
	benchmark()
	<-stopCh
	zaplogger.Sugar().Info("game/server trigger shutdown")
	<-stopCh
	zaplogger.Sugar().Info("game/server shutdown gracefully")
}

func benchmark() {
	tag := "TestClient"
	opt, err := rocketmqrestapi.GetOptionFromEnv()
	if err != nil {
		panic(err)
	}
	client := rocketmqrestapi.NewClient(opt, "unit-test")

	var msg mq_http_sdk.PublishMessageRequest
	msg = mq_http_sdk.PublishMessageRequest{
		MessageBody: "hello mq!xxxxxxxxxsasjamslaklwqmawds291u90msalsma29003isam21s121slw1-2mi1-21saks;a;sas;asasls" +
			"s,a3845ufjdo39djk30d30dkdpdkd-d-23ikd-23kd-",
		Properties: map[string]string{
			"pid": "123",
		},
	}
	//var wg sync.WaitGroup
	//wg.Add(100)
	for i := 0; i < 100; i++ {
		//go func(i int) {
			msg.Properties["pid"] = strconv.Itoa(i)
			t1 := time.Now()
			ret, err := client.Publish(tag, msg)
			fmt.Printf("used in ms:%d\n", time.Now().Sub(t1).Milliseconds())
			if err == nil {
				fmt.Printf("Publish ---->\n\tMessageId:%s, BodyMD5:%s, \n", ret.MessageId, ret.MessageBodyMD5)
			}
			//wg.Done()
		//}(i)
	}
	//wg.Wait()
}
