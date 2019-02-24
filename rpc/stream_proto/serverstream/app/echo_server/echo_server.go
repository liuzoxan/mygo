package main

import (
	"context"
	. "github.com/lzxbill7/mygo/rpc/stream_proto/serverstream"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	port = ":39001"
)

type server struct {
}

func (*server) Echo(req *EchoReq, echoServer Echo_EchoServer) error {
	log.Println("Echo: ", req.Msg)
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 3)
	for {
		select {
		case <- ctx.Done():
			return nil
		default:
			echoServer.Send(&EchoRes{Msg:"pong"})
		}
	}
}

func main() {
	s := grpc.NewServer()
	RegisterEchoServer(s, &server{})
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}else{
		log.Printf("server listen : %v", port)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
