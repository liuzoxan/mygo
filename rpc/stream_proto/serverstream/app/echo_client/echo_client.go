package main

import (
	"context"
	. "github.com/lzxbill7/mygo/rpc/stream_proto/serverstream"
	"google.golang.org/grpc"
	"io"
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
	for {
		data, err := r.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("server error %v", err)
		}
		log.Println("from server", data)
	}
}
