SHELL := $(shell which bash) # set default shell
# OS / Arch we will build our binaries for
OSARCH := "linux/amd64 linux/386 windows/amd64 windows/386 darwin/amd64 darwin/386"
ENV = /usr/bin/env

.SILENT: ; # no need for @
.ONESHELL: ; # recipes execute in same shell
.NOTPARALLEL: ; # wait for this target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

.PHONY: all # All targets are accessible for user
.DEFAULT: help # Running Make will run the help target

help: ## Show Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

dep: ## Get build dependencies
	rm -rf vendor/ && go get -v -u github.com/golang/dep/cmd/dep && dep ensure -v

config: ## Config the app
	cp .env.example.json .env.json

build: ## Build the app
	go install && cp .env.json $(GOPATH)/bin/

test: ## Launch tests
	cp .env.test.example.json .env.test.json && go test -v ./...

all: ## Install, Test and build
	make dep && make test && make config && make build
