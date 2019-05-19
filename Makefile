# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=urlinsane
VERSION=$(shell grep -e 'VERSION = ".*"' urlinsane.go | cut -d= -f2 | sed  s/[[:space:]]*\"//g)

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'

build: ## Build the binaries for Windows, OSX, and Linux
build: deps
	mkdir -p builds
	cd cmd; $(GOBUILD) -o ../builds/$(BINARY_NAME) -v
	cd cmd; env GOOS=darwin GOARCH=amd64 $(GOBUILD) -o ../builds/$(BINARY_NAME)-$(VERSION)-darwin-amd64 -v
	cd cmd; env GOOS=linux GOARCH=amd64 $(GOBUILD) -o ../builds/$(BINARY_NAME)-$(VERSION)-linux-amd64 -v
	cd cmd; env GOOS=windows GOARCH=amd64 $(GOBUILD) -o ../builds/$(BINARY_NAME)-$(VERSION)-windows-amd64.exe -v
	md5 builds/$(BINARY_NAME) | md5sum builds/$(BINARY_NAME)

deps: ## Install dependencies
	$(GOGET) ./...
	# $(GOGET) github.com/rangertaha/urlinsane

docker: image ## Build docker image and upload to docker hub
	docker login

image: clean ## Build docker image
	docker build -t $(BINARY_NAME) .

test: ## Run unit test
test: deps
	$(GOTEST) -v ./...

clean: ## Remove files created by the build
	$(GOCLEAN)
	rm -fr builds

run: ## Run server docker image
	docker run -it --rm -p 8080:8080 urlinsane
