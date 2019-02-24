package main

import (
	. "github.com/lzxbill7/mygo/rpc/stream_proto/clientstream"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

const (
	port = ":39001"
)

type server struct {
}

func (s *server) Echo(echoServer Echo_EchoServer) error {
	for {
		req, err := echoServer.Recv()
		if err == io.EOF {
			break
		}
		log.Printf("from client %v", req.Msg)
	}
	echoServer.SendAndClose(&EchoRes{Msg:"pong"})
	return nil
}

func main()  {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to lis %v", err)
	}
	s := grpc.NewServer()
	RegisterEchoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
