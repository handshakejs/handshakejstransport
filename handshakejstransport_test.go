package handshakejstransport_test

import (
	"../handshakejstransport"
	"testing"
)

func TestSetup(t *testing.T) {
	handshakejstransport.Setup("smtp.sendgrid.net", "587", "username", "password")
}
