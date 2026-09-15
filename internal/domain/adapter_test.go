package domain

import "testing"

func TestAdapterFor(t *testing.T) {
	t.Parallel()
	cases := []struct {
		fqdn    string
		tld     string
		channel Channel
	}{
		{"google.com", "com", ChannelRDAP},
		{"GOOGLE.NET", "net", ChannelRDAP},
		{"example.org", "org", ChannelRDAP},
		{"google.dev", "dev", ChannelRDAP},
		{"foo.app", "app", ChannelRDAP},
		{"bar.xyz", "xyz", ChannelRDAP},
		{"x.info", "info", ChannelRDAP},
		{"y.ai", "ai", ChannelRDAP},
		{"z.cc", "cc", ChannelRDAP},
		{"a.fun", "fun", ChannelRDAP},
		{"a.site", "site", ChannelRDAP},
		{"a.online", "online", ChannelRDAP},
		{"a.store", "store", ChannelRDAP},
		{"a.tech", "tech", ChannelRDAP},
		{"a.pw", "pw", ChannelRDAP},
		{"zzzfox.cn", "cn", ChannelWhois},
		{"google.io", "io", ChannelWhois},
		{"nic.me", "me", ChannelWhois},
		{"google.co", "co", ChannelWhois},
	}
	for _, tc := range cases {
		t.Run(tc.fqdn, func(t *testing.T) {
			t.Parallel()
			a, err := AdapterFor(tc.fqdn)
			if err != nil {
				t.Fatalf("AdapterFor(%q): %v", tc.fqdn, err)
			}
			if a.TLD != tc.tld {
				t.Fatalf("tld = %q, want %q", a.TLD, tc.tld)
			}
			if a.Channel != tc.channel {
				t.Fatalf("channel = %q, want %q", a.Channel, tc.channel)
			}
			if a.Channel == ChannelRDAP && a.RDAPURL == "" {
				t.Fatalf("rdap adapter missing URL")
			}
			if a.Channel == ChannelWhois && a.WhoisHost == "" {
				t.Fatalf("whois adapter missing host")
			}
		})
	}
}

func TestAdapterForUnsupported(t *testing.T) {
	t.Parallel()
	_, err := AdapterFor("example.de")
	if err == nil {
		t.Fatal("expected error for unsupported tld")
	}
}

func TestAdapterForInvalid(t *testing.T) {
	t.Parallel()
	for _, fqdn := range []string{"", "nodot", ".com", "a."} {
		if _, err := AdapterFor(fqdn); err == nil {
			t.Fatalf("expected error for %q", fqdn)
		}
	}
}

func TestSupportedContainsV1List(t *testing.T) {
	t.Parallel()
	want := []string{
		"com", "net", "org", "io", "me", "co",
		"dev", "app", "xyz", "info", "ai", "cc",
		"fun", "site", "online", "store", "tech", "pw", "cn",
	}
	got := SupportedTLDs()
	if len(got) != len(want) {
		t.Fatalf("SupportedTLDs len = %d, want %d: %v", len(got), len(want), got)
	}
	set := map[string]bool{}
	for _, tld := range got {
		set[tld] = true
	}
	for _, tld := range want {
		if !set[tld] {
			t.Fatalf("missing tld %q", tld)
		}
	}
}

func TestNormalizeFQDN(t *testing.T) {
	t.Parallel()
	got, err := NormalizeFQDN("  Google.COM. ")
	if err != nil {
		t.Fatal(err)
	}
	if got != "google.com" {
		t.Fatalf("got %q", got)
	}
}
