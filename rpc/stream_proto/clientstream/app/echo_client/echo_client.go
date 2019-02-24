package main

import (
	"context"
	. "github.com/lzxbill7/mygo/rpc/stream_proto/clientstream"
	"google.golang.org/grpc"
	"io"
	"log"
)

const (
	addr = "localhost:39001"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dail %v", err)
	}
	defer conn.Close()
	c := NewEchoClient(conn)
	var ctx = context.Background()
	client, err := c.Echo(ctx)
	if err != nil {
		log.Fatalf("failed to Echo")
	}
	client.Send(&EchoReq{Msg:"ping1"})
	client.Send(&EchoReq{Msg:"ping2"})
	client.Send(&EchoReq{Msg:"ping3"})
	res, err := client.CloseAndRecv()
	if err == io.EOF {
		log.Printf("from server end")
	} else if err != nil {
		log.Fatalf("failed to recv %v", err)
	} else {
		log.Printf("from server %v", res.Msg)
	}
}
