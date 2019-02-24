# set compiler, protoc use protoc-gen-go plugin, please make sure it!
PROTOC=~/go/bin/protoc/bin/protoc
GOCMD=go

all: server client codegen

server:
	-cd app/echo_server && ${GOCMD} build echo_server.go
.PHONY: server

client:
	-cd app/echo_client && ${GOCMD} build echo_client.go
.PHONY: client

codegen: echo.pb.go
echo.pb.go: echo.proto
	-${PROTOC} --go_out=plugins=grpc:. echo.proto

clean:
	-rm -rf echo.pb.go
