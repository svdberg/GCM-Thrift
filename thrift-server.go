package main

import (
	"fmt"

	"github.com/svdberg/gopush/gen-go/gopush"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) error {
	var transport thrift.TServerTransport
	var err error
	transport, err = thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}

	fmt.Printf("%T\n", transport)
	handler := NewGopushHandler()
	processor := gopush.NewGopushProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", addr)
	return server.Serve()
}
