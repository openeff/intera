package mcp

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type StatusInput struct {
	App string `json:"app" jsonschema:"Application to inspect"`
}

type StatusOutput struct {
	App    string `json:"app"`
	Status string `json:"status"`
}

func GetStatus(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input StatusInput,
) (*mcp.CallToolResult, StatusOutput, error) {
	return nil, StatusOutput{
		App:    input.App,
		Status: "connected",
	}, nil
}
