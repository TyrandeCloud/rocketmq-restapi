.PHONY: build mod

GO_ROOT := $(shell cd && pwd)
HARBOR_DOMAIN := $(shell echo ${HARBOR})

mod:
	go mod download
	go mod tidy