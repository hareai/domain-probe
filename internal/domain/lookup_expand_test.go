package domain

import "testing"

func TestExpand(t *testing.T) {
	got := Expand(" GoFox ")
	want := len(SupportedTLDs())
	if len(got) != want {
		t.Fatalf("len %d want %d", len(got), want)
	}
	if got[0] != "gofox.com" {
		t.Fatalf("first %q", got[0])
	}
	last := got[len(got)-1]
	if last != "gofox.cn" {
		t.Fatalf("last %q", last)
	}
}
