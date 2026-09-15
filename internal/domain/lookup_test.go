package domain

import (
	"context"
	"errors"
	"testing"
)

type fakeGet struct {
	rdap  map[string]int
	whois map[string]string
	err   error
}

func (f fakeGet) RDAP(_ context.Context, url string) (int, string, error) {
	if f.err != nil {
		return 0, "", f.err
	}
	code, ok := f.rdap[url]
	if !ok {
		return 0, "", errors.New("no stub for " + url)
	}
	return code, "", nil
}

func (f fakeGet) Whois(_ context.Context, _, fqdn string) (string, error) {
	if f.err != nil {
		return "", f.err
	}
	text, ok := f.whois[fqdn]
	if !ok {
		return "", errors.New("no stub for " + fqdn)
	}
	return text, nil
}

func TestLookupRDAPTaken(t *testing.T) {
	p := Probe{Get: fakeGet{rdap: map[string]int{
		"https://rdap.verisign.com/com/v1/domain/google.com": 200,
	}}}
	got, err := p.Lookup(context.Background(), "google.com")
	if err != nil {
		t.Fatal(err)
	}
	if got.Status != StatusRegistered || got.Channel != ChannelRDAP {
		t.Fatalf("%+v", got)
	}
}

func TestLookupRDAPEmpty(t *testing.T) {
	p := Probe{Get: fakeGet{rdap: map[string]int{
		"https://rdap.verisign.com/com/v1/domain/zzzzqzzzx.com": 404,
	}}}
	got, err := p.Lookup(context.Background(), "zzzzqzzzx.com")
	if err != nil {
		t.Fatal(err)
	}
	if got.Status != StatusAvailable {
		t.Fatalf("%+v", got)
	}
}

func TestLookupWhoisCN(t *testing.T) {
	p := Probe{Get: fakeGet{whois: map[string]string{
		"zzzfox.cn": "No matching record.\n",
	}}}
	got, err := p.Lookup(context.Background(), "zzzfox.cn")
	if err != nil {
		t.Fatal(err)
	}
	if got.Status != StatusAvailable || got.Channel != ChannelWhois {
		t.Fatalf("%+v", got)
	}
}

func TestLookupUnsupported(t *testing.T) {
	p := Probe{Get: fakeGet{}}
	_, err := p.Lookup(context.Background(), "example.de")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestLookupGetErrorIsUnknown(t *testing.T) {
	p := Probe{Get: fakeGet{err: errors.New("timeout")}}
	got, err := p.Lookup(context.Background(), "google.com")
	if err != nil {
		t.Fatal(err)
	}
	if got.Status != StatusUnknown {
		t.Fatalf("%+v", got)
	}
}

func TestLookupMany(t *testing.T) {
	p := Probe{Get: fakeGet{
		rdap: map[string]int{
			"https://rdap.verisign.com/com/v1/domain/google.com": 200,
			"https://rdap.verisign.com/com/v1/domain/zzzzqzzzx.com": 404,
		},
	}}
	got := p.LookupMany(context.Background(), []string{"google.com", "zzzzqzzzx.com"})
	if len(got) != 2 {
		t.Fatalf("len %d", len(got))
	}
	if got[0].Status != StatusRegistered || got[1].Status != StatusAvailable {
		t.Fatalf("%+v", got)
	}
}
