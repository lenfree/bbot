DEFAULT_NAME := bbot

package = github.com/lenfree/$(DEFAULT_NAME)

all: release

.PHONY: install
install:
	go get -v

.PHONY: script
script:
	go vet ./...
	go build -v -race ./...

.PHONY: release
release: install script
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/$(DEFAULT_NAME)-linux-amd64 $(package)
	GOOS=linux GOARCH=arm go build -o release/$(DEFAULT_NAME)-linux-amd64 $(package)
	GOOS=darwin GOARCH=amd64 go build -o release/$(DEFAULT_NAME)-darwin-amd64 $(package)
