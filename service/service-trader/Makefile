
GOPATH:=$(shell go env GOPATH)

build:
	CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o trader-srv main.go