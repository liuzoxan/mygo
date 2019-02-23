package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	. "github.com/lzxbill7/mygo/rpc/echo_thrift/gen-go/echo_thrift"
	"log"
	"os"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of", os.Args[0])
	flag.PrintDefaults()
}

func handleClient(client *EchoClient) (err error) {
	res, err := client.Echo(&Req{Msg:"Hello!"})
	if err != nil {
		return err
	}
	log.Println("echo :", res.Msg)
	return nil
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) error {
	var transport thrift.TTransport
	var err error
	if secure {
		// see: same as server
	}
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err
	}
	transport= transportFactory.GetTransport(transport)
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
		}
	return handleClient(NewEchoClientFactory(transport, protocolFactory))

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
	if err := runClient(transportFactory, protocolFactory, *addr, *secure); err != nil {
		log.Fatalf("failed to client: %v", err)
	}
}
