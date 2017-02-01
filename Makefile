NAME     := gpv
VERSION  := v0.0.1
REVISION := $(shell git rev-parse --short HEAD)

SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""

.PHONY: dep
dep:
	go get -u github.com/golang/dep/...

.PHONY: dep-install
dep-install: dep
	dep ensure

.PHONY: dep-init
dep-init: dep
	dep init

.PHONY: install
install:
	go install $(LDFLAGS)

.PHONY: test
test:
	go test -cover -v ../...

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf vendor/*

