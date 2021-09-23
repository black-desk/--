package main

import (
	"fmt"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
	"github.com/tliron/kutil/logging"

	// Must include a backend implementation. See kutil's logging/ for other options.
	_ "github.com/tliron/kutil/logging/simple"
)

const lsName = "my language"

var version string = "0.0.1"
var handler protocol.Handler

var fileMap map[protocol.URI]string

func init() {
	fileMap = map[protocol.URI]string{}
}

func main() {
	// Must configure logging before using server. Note that the import of a backend
	// is required as well.
	// See kutil's logging.Backend for Configure() param details.
	logging.Configure(1, nil)

	handler = protocol.Handler{
		Initialize:        initialize,
		Initialized:       initialized,
		Shutdown:          shutdown,
		SetTrace:          setTrace,
		TextDocumentHover: textDocumentHover,
	}

	server := server.NewServer(&handler, lsName, false)

	server.RunStdio()
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (interface{}, error) {
	capabilities := handler.CreateServerCapabilities()

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lsName,
			Version: &version,
		},
	}, nil
}

func initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	return nil
}

func shutdown(context *glsp.Context) error {
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func setTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}

func textDocumentHover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	return &protocol.Hover{
		Contents: &protocol.MarkupContent{
			Kind:  protocol.MarkupKindPlainText,
			Value: fmt.Sprint("Character: ", params.Position.Character),
		},
		Range: &protocol.Range{
			Start: params.Position,
			End:   params.Position,
		},
	}, nil
}
