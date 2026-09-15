package main

import (
	"context"
	"log"
	"time"

	"github.com/hareai/domain-probe/internal/domain"
	"github.com/hareai/domain-probe/internal/server"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	p := domain.Probe{Get: domain.Net{}, Gap: 2 * time.Second}
	s := server.New(p)
	if err := s.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
