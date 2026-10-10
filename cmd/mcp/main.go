package main

import (
	"context"
	"log"

	m "openeff/intera/internal/mcp"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "intera",
			Version: "0.0.0",
		},
		nil,
	)

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "get_app_status",
			Description: "Get the connection status of an application",
		},
		m.GetStatus,
	)

	if err := server.Run(
		context.Background(),
		&mcp.StdioTransport{},
	); err != nil {
		log.Fatal(err)
	}
}
