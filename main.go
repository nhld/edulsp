package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"edulsp/analysis"
	"edulsp/lsp"
	"edulsp/rpc"
)

func main() {
	fmt.Println("edulsp is running...")
	logger := getLogger("./log.txt")
	logger.Println("edulsp started.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodingMessage(msg)
		if err != nil {
			logger.Printf("Error: %s", err)
			continue
		}
		handleMessage(logger, writer, state, method, contents)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, contents []byte) {
	logger.Printf("Received method: %s", method)
	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Cant parse: %s", err)
		}
		logger.Printf("Connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

		msg := lsp.NewIninitializeResponse(request.ID)
		err := writeResponse(writer, msg)
		if err != nil {
			log.Printf("Failed to write response: %v", err)
		}
		logger.Print("Sent reply")

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didOpen: %s", err)
			return
		}
		logger.Printf("Opened: %s", request.Params.TextDocument.URI)
		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)

	case "textDocument/didChange":
		var request lsp.TextDocumentDidChangeNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didChange: %s", err)
			return
		}
		logger.Printf("Changed: %s", request.Params.TextDocument.URI)
		for _, change := range request.Params.ContentChanges {
			state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		}

	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/hover: %s", err)
			return
		}
		reponse := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		err := writeResponse(writer, reponse)
		if err != nil {
			log.Printf("Failed to write response: %v", err)
		}
	}
}

func writeResponse(writer io.Writer, msg any) error {
	reply := rpc.EncodingMessage(msg)
	_, err := writer.Write([]byte(reply))
	if err != nil {
		return err
	}
	return nil
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("No log file!")
	}
	return log.New(logfile, "[edulsp] ", log.Ldate|log.Ltime|log.Lshortfile)
}
