package rpc_test

import (
	"testing"

	"edulsp/rpc"
)

type EncodingTesting struct {
	Testing bool
}

func TestEncodeMessage(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodingMessage(EncodingTesting{Testing: true})
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestDecodeMessage(t *testing.T) {
	incommingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodingMessage([]byte(incommingMessage))
	contentLength := len(content)

	if err != nil {
		t.Fatal(err)
	}
	if contentLength != 15 {
		t.Fatalf("Expected: 15, Actual: %d", contentLength)
	}
	if method != "hi" {
		t.Fatalf("Expected: hi, Actual: %s", method)
	}
}
