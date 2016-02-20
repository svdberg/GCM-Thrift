package main

import (
	"github.com/google/go-gcm"
	m "github.com/svdberg/gopush/gen-go/gopush"
)

type GopushHandler struct {
}

func NewGopushHandler() *GopushHandler {
	return &GopushHandler{}
}

func (p *GopushHandler) Message(request *m.GcmRequest) error {
	return sendMessage()
}

// Handle request to send a new message.
func sendMessage() error {

	const apiKey = "12213231132312"

	// Send HTTP message
	var m gcm.HttpMessage = gcm.HttpMessage{}
	var n gcm.Notification = gcm.Notification{}
	n.Title = "Portugal vs. Denmark"
	n.Body = "5 to 1"

	m.To = "1231323113"
	m.Notification = n

	_, sendErr := gcm.SendHttp(apiKey, m)
	return sendErr
}
