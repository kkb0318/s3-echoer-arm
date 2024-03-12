.PHONY: build

build:
	GOOS=linux GOARCH=arm64 go build -o bin/s3-echoer-linux-arm64 .
