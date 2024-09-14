package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"edulsp/lsp"
	"edulsp/rpc"
)

func main() {
	fmt.Println("edulsp is running...")
	logger := getLogger("./log.txt")
	logger.Println("edulsp started.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodingMessage(msg)
		if err != nil {
			logger.Printf("Error: %s", err)
			continue
		}
		handleMessage(logger, method, contents)
	}
}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("Received method: %s", method)
	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Cant parse: %s", err)
		}
		logger.Printf("Connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("No log file!")
	}
	return log.New(logfile, "[edulsp] ", log.Ldate|log.Ltime|log.Lshortfile)
}
