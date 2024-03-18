package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodingMessage(msg any) string {
	bytesMsg, err := json.Marshal(msg)
	if err != nil {
		panic("This should never happen")
	}
	length := len(bytesMsg)
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", length, bytesMsg)
}

type BaseMessage struct {
	Method string `json:"method"`
}

func DecodingMessage(msg []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("Could not find the seperator")
	}

	contentLengthStr := string(header[len("Content-Length: "):])
	contentLength, err := strconv.Atoi(contentLengthStr)
	if err != nil {
		return "", nil, err
	}

	var baseMessage BaseMessage
	if err := json.Unmarshal(content, &baseMessage); err != nil {
		return "", nil, err
	}

	return baseMessage.Method, content[:contentLength], nil
}
