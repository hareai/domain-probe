//go:build live

package domain

import (
	"context"
	"testing"
	"time"
)

func TestLiveLookup(t *testing.T) {
	p := Probe{Get: Net{}}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	got, err := p.Lookup(ctx, "google.com")
	if err != nil {
		t.Fatal(err)
	}
	if got.Status != StatusRegistered {
		t.Fatalf("google.com: %+v", got)
	}

	got, err = p.Lookup(ctx, "zzzfox.cn")
	if err != nil {
		t.Fatal(err)
	}
	if got.Status != StatusAvailable {
		t.Fatalf("zzzfox.cn: %+v", got)
	}
}
