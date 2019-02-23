package main

import (
	"context"
	. "github.com/lzxbill7/mygo/rpc/echo_proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:39001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	c := NewEchoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Echo(ctx, &EchoReq{Msg:"Hello!"})
	if err != nil {
		log.Fatalf("failed to Echo server: %v", err)
	}
	log.Println("echo from server: ", r.Msg)
}
