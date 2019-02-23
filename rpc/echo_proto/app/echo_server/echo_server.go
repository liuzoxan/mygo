package main

import (
	"context"
	. "github.com/lzxbill7/mygo/rpc/echo_proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":39001"
)

type server struct {
}

func (*server) Echo(ctx context.Context, req *EchoReq) (*EchoRes, error) {
	log.Println("Echo: ", req.Msg)
	return &EchoRes{Msg: "Echoed"}, nil
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
