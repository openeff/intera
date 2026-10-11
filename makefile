MCP_GEN_TOOL := $(CURDIR)/build/mcp_intera
CLI_TOOL := $(CURDIR)/build/intera


.PHONY: build-mcp build-cli

build-mcp:
	@echo "Building mcp_gen_tool..."
	@mkdir -p "$(CURDIR)/build"
	@go build -o "$(MCP_GEN_TOOL)" ./cmd/mcp

build-cli:
	@echo "Building intera..."
	@mkdir -p "$(CURDIR)/build"
	@go build -o "$(CLI_TOOL)" ./cmd/