package domain

import "testing"

func TestClassifyWhoisCN(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		fqdn string
		text string
		want Status
	}{
		{
			name: "empty NMR only",
			fqdn: "zzzfox.cn",
			text: "No matching record.\n",
			want: StatusAvailable,
		},
		{
			name: "empty NMR after querying banner",
			fqdn: "zzzfox.cn",
			text: "[Querying whois.cnnic.cn]\n[whois.cnnic.cn]\nNo matching record.\n",
			want: StatusAvailable,
		},
		{
			name: "taken",
			fqdn: "google.cn",
			text: "Domain Name: google.cn\nROID: 20030311s10001s00024537-cn\nDomain Status: ok\n",
			want: StatusRegistered,
		},
		{
			name: "pending delete",
			fqdn: "example.cn",
			text: "Domain Name: example.cn\nDomain Status: pendingDelete\nDomain Status: inactive\n",
			want: StatusPendingDelete,
		},
		{
			name: "restricted",
			fqdn: "nfox.cn",
			text: "the Domain Name you apply can not be registered online. Please consult your Domain Name registrar\n",
			want: StatusReserved,
		},
		{
			name: "ambiguous NMR plus domain name",
			fqdn: "example.cn",
			text: "No matching record.\nDomain Name: example.cn\n",
			want: StatusUnknown,
		},
		{
			name: "empty body",
			fqdn: "example.cn",
			text: "",
			want: StatusUnknown,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := ClassifyWhois("cn", tc.fqdn, tc.text)
			if got != tc.want {
				t.Fatalf("ClassifyWhois(cn, %q) = %q, want %q", tc.fqdn, got, tc.want)
			}
		})
	}
}

func TestClassifyWhoisIO(t *testing.T) {
	t.Parallel()
	if got := ClassifyWhois("io", "zzzqwxio.io", "Domain not found.\n>>> Last update of WHOIS database: 2026-09-15T05:25:36Z <<<\n"); got != StatusAvailable {
		t.Fatalf("empty io: got %q", got)
	}
	if got := ClassifyWhois("io", "google.io", "Domain Name: google.io\nRegistrar: MarkMonitor Inc.\n"); got != StatusRegistered {
		t.Fatalf("taken io: got %q", got)
	}
}

func TestClassifyWhoisME(t *testing.T) {
	t.Parallel()
	if got := ClassifyWhois("me", "zzzqwxme.me", "Domain not found.\n"); got != StatusAvailable {
		t.Fatalf("empty me: got %q", got)
	}
	if got := ClassifyWhois("me", "nic.me", "Domain Name: nic.me\n"); got != StatusRegistered {
		t.Fatalf("taken me: got %q", got)
	}
}

func TestClassifyWhoisCO(t *testing.T) {
	t.Parallel()
	if got := ClassifyWhois("co", "zzzqwxco.co", "The queried object does not exist: DOMAIN NOT FOUND\n\n>>> Last update of WHOIS database: 2026-09-15T05:26:58.0Z <<<\n"); got != StatusAvailable {
		t.Fatalf("empty co: got %q", got)
	}
	if got := ClassifyWhois("co", "google.co", "Domain Name: GOOGLE.CO\nRegistry Domain ID: D157997-CNIC\n"); got != StatusRegistered {
		t.Fatalf("taken co: got %q", got)
	}
}

func TestClassifyRDAP(t *testing.T) {
	t.Parallel()
	cases := []struct {
		code int
		want Status
	}{
		{200, StatusRegistered},
		{404, StatusAvailable},
		{403, StatusUnknown},
		{429, StatusUnknown},
		{500, StatusUnknown},
		{0, StatusUnknown},
	}
	for _, tc := range cases {
		got := ClassifyRDAP(tc.code)
		if got != tc.want {
			t.Fatalf("ClassifyRDAP(%d) = %q, want %q", tc.code, got, tc.want)
		}
	}
}
