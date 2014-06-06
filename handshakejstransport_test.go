package handshakejstransport_test

import (
	"../handshakejstransport"
	"testing"
)

func TestSetup(t *testing.T) {
	options := &handshakejstransport.Options{"smtp.sendgrid.net", "587", "username", "password"}
	handshakejstransport.Setup(options)
}
