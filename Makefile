.PHONY:mod build

GO_ROOT := $(shell cd && pwd)
HARBOR_DOMAIN := $(shell echo ${HARBOR})
PROJECT := hamster
IMAGE := "$(HARBOR_DOMAIN)/$(PROJECT)/rocketmq-test:latest"

mod:
	go mod download
	go mod tidy

build:
	-i docker image rm $(IMAGE)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o rocketmq cmd/main.go
	docker build -t $(IMAGE) .
	rm -f rocketmq
	docker push $(IMAGE)

run:
	go run cmd/main.go