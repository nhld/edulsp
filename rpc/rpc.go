package rpc

import (
	"encoding/json"
	"fmt"
)

func EncodingMessage(msg any) string {
	bytesMsg, err := json.Marshal(msg)
	if err != nil {
		panic("This should never happen")
	}
	length := len(bytesMsg)
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", length, bytesMsg)
}
