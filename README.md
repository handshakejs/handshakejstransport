# handshakejstransport

<img src="https://raw.githubusercontent.com/scottmotte/handshakejstransport/master/handshakejslogictransport.gif" alt="handshakejstransport" align="right" />

[![BuildStatus](https://travis-ci.org/handshakejs/handshakejstransport.png?branch=master)](https://travis-ci.org/handshakejs/handshakejstransport)

Transport mechanisms for delivering handshakejs authcodes to people.

This library is part of the larger [Handshake.js ecosystem](https://github.com/handshakejs).

## Usage

```go
package main

import (
  "fmt"
  handshakejstransport "github.com/handshakejs/handshakejstransport"
)

func main() {
  handshakejstransport.Setup("smtp.sendgrid.net", "587", "username", "password")
  handshakejstransport.ViaEmail("person0@mailinator.com", "Your authcode is 1234", "This is the text of the email", "This is the <b>html</b> of the email")
}
```

## Installation

```
go get github.com/handshakejs/handshakejstransport
```

## Running Tests

```
go test -v
```
