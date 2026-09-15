package server

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hareai/domain-probe/internal/domain"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type fakeGet struct {
	rdap map[string]int
}

func (f fakeGet) RDAP(_ context.Context, url string) (int, string, error) {
	return f.rdap[url], "", nil
}

func (f fakeGet) Whois(context.Context, string, string) (string, error) {
	return "", nil
}

func TestTools(t *testing.T) {
	p := domain.Probe{Get: fakeGet{rdap: map[string]int{
		"https://rdap.verisign.com/com/v1/domain/google.com":    200,
		"https://rdap.verisign.com/com/v1/domain/zzzzqzzzx.com": 404,
	}}}
	s := New(p)
	c := mcp.NewClient(&mcp.Implementation{Name: "test", Version: "v0"}, nil)
	t1, t2 := mcp.NewInMemoryTransports()
	if _, err := s.Connect(context.Background(), t1, nil); err != nil {
		t.Fatal(err)
	}
	cs, err := c.Connect(context.Background(), t2, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { cs.Close() })

	var names []string
	for tool, err := range cs.Tools(context.Background(), nil) {
		if err != nil {
			t.Fatal(err)
		}
		names = append(names, tool.Name)
	}
	if len(names) != 2 {
		t.Fatalf("tools %v", names)
	}

	res, err := cs.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "lookup",
		Arguments: map[string]any{"name": "google.com"},
	})
	if err != nil {
		t.Fatal(err)
	}
	raw, _ := json.Marshal(res.StructuredContent)
	var rec domain.Record
	if err := json.Unmarshal(raw, &rec); err != nil {
		t.Fatal(err, string(raw))
	}
	if rec.Status != domain.StatusRegistered {
		t.Fatalf("%+v", rec)
	}

	res, err = cs.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "lookup_batch",
		Arguments: map[string]any{"names": []string{"google.com", "zzzzqzzzx.com"}},
	})
	if err != nil {
		t.Fatal(err)
	}
	raw, _ = json.Marshal(res.StructuredContent)
	var out batchOut
	if err := json.Unmarshal(raw, &out); err != nil {
		t.Fatal(err, string(raw))
	}
	if len(out.Results) != 2 || out.Results[1].Status != domain.StatusAvailable {
		t.Fatalf("%+v", out)
	}
}
