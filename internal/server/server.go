package server

import (
	"context"

	"github.com/hareai/domain-probe/internal/domain"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func New(p domain.Probe) *mcp.Server {
	s := mcp.NewServer(&mcp.Implementation{Name: "domain-probe", Version: "0.1.0"}, nil)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "lookup",
		Description: "Look up one domain name.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, in lookupIn) (*mcp.CallToolResult, domain.Record, error) {
		rec, err := p.Lookup(ctx, in.Name)
		return nil, rec, err
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "lookup_batch",
		Description: "Look up several domain names, or one label across the v1 suffixes.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, in batchIn) (*mcp.CallToolResult, batchOut, error) {
		names := in.Names
		if in.Label != "" {
			names = domain.Expand(in.Label)
		}
		return nil, batchOut{Results: p.LookupMany(ctx, names)}, nil
	})

	return s
}

type lookupIn struct {
	Name string `json:"name" jsonschema:"the domain name, for example google.com"`
}

type batchIn struct {
	Names []string `json:"names,omitempty" jsonschema:"fully qualified names to look up"`
	Label string   `json:"label,omitempty" jsonschema:"bare label; looked up on every v1 suffix"`
}

type batchOut struct {
	Results []domain.Record `json:"results"`
}
