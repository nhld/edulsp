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
