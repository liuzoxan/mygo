package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	. "github.com/lzxbill7/mygo/rpc/echo_thrift/gen-go/echo_thrift"
	"log"
	"os"
)

type EchoHandler struct {
}

func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

func (handler *EchoHandler) Echo(req *Req) (*Res, error) {
	return &Res{Msg: req.Msg}, nil
}

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) error {
	if secure {
		// for secure see: http://thrift.apache.org/tutorial/go
	}
	var transport thrift.TServerTransport
	var err error
	transport, err = thrift.NewTServerSocket(addr)

	if err != nil {
		return err
	}
	log.Printf("%T %T", transportFactory, protocolFactory)
	handler := NewEchoHandler()
	processor := NewEchoProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	log.Println("Starting the simple server... on ", addr)
	return server.Serve()
}

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage
	protocol := flag.String("P", "binary", "Specify the protocol (binary, compact, json, simplejson)")
	framed := flag.Bool("framed", false, "Use framed transport")
	buffered := flag.Bool("buffered", false, "Use buffered transport")
	addr := flag.String("addr", "localhost:39002", "Address to listen to")
	secure := flag.Bool("secure", false, "Use tls secure transport")
	flag.Parse()

	// protocol
	var protocolFactory thrift.TProtocolFactory
	switch *protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		fmt.Fprint(os.Stderr, "Invalid protocol specified", protocol)
		Usage()
		os.Exit(1)
	}

	// transport
	var transportFactory thrift.TTransportFactory
	if *buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}
	if *framed {
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}

	// run
	if err := runServer(transportFactory, protocolFactory, *addr, *secure); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
