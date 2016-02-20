package main

import (
	"log"

	gopush "github.com/svdberg/gopush/gen-go/gopush"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func main() {
	const thrifthost = "localhost:1337"
	s, err := thrift.NewTSocket(thrifthost)
	if err != nil {
		log.Fatalf("Error connecting to the elastic-indexer-service: %s", err)
	}

	t := thrift.NewTBufferedTransport(s, 1024)
	pf := thrift.NewTBinaryProtocolFactoryDefault()
	gopushClient := gopush.NewGopushClientFactory(t, pf)
	if err := t.Open(); err != nil {
		log.Fatalf("Error connecting to go push stuff whadajeah: %s", err)
	}
	defer t.Close()

	m := gopush.NewGcmRequest()
	payload := gopush.NewPayload()
	payload.Message = "My payload that I will never see"
	m.Data = payload

	gopushClient.Message(m)
}
