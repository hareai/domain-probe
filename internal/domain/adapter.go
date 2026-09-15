package domain

import (
	"fmt"
	"strings"
)

// Adapter describes how to look up one TLD. v1 only covers the list in adapters.
type Adapter struct {
	TLD       string
	Channel   Channel
	RDAPURL   string // prefix; caller appends the FQDN
	WhoisHost string
}

var adapters = map[string]Adapter{
	"com":    {TLD: "com", Channel: ChannelRDAP, RDAPURL: "https://rdap.verisign.com/com/v1/domain/"},
	"net":    {TLD: "net", Channel: ChannelRDAP, RDAPURL: "https://rdap.verisign.com/net/v1/domain/"},
	"org":    {TLD: "org", Channel: ChannelRDAP, RDAPURL: "https://rdap.publicinterestregistry.org/rdap/domain/"},
	"dev":    {TLD: "dev", Channel: ChannelRDAP, RDAPURL: "https://pubapi.registry.google/rdap/domain/"},
	"app":    {TLD: "app", Channel: ChannelRDAP, RDAPURL: "https://pubapi.registry.google/rdap/domain/"},
	"xyz":    {TLD: "xyz", Channel: ChannelRDAP, RDAPURL: "https://rdap.centralnic.com/xyz/domain/"},
	"info":   {TLD: "info", Channel: ChannelRDAP, RDAPURL: "https://rdap.identitydigital.services/rdap/domain/"},
	"ai":     {TLD: "ai", Channel: ChannelRDAP, RDAPURL: "https://rdap.identitydigital.services/rdap/domain/"},
	"cc":     {TLD: "cc", Channel: ChannelRDAP, RDAPURL: "https://tld-rdap.verisign.com/cc/v1/domain/"},
	"fun":    {TLD: "fun", Channel: ChannelRDAP, RDAPURL: "https://rdap.radix.host/rdap/domain/"},
	"site":   {TLD: "site", Channel: ChannelRDAP, RDAPURL: "https://rdap.radix.host/rdap/domain/"},
	"online": {TLD: "online", Channel: ChannelRDAP, RDAPURL: "https://rdap.radix.host/rdap/domain/"},
	"store":  {TLD: "store", Channel: ChannelRDAP, RDAPURL: "https://rdap.radix.host/rdap/domain/"},
	"tech":   {TLD: "tech", Channel: ChannelRDAP, RDAPURL: "https://rdap.radix.host/rdap/domain/"},
	"pw":     {TLD: "pw", Channel: ChannelRDAP, RDAPURL: "https://rdap.radix.host/rdap/domain/"},
	"cn":     {TLD: "cn", Channel: ChannelWhois, WhoisHost: "whois.cnnic.cn"},
	"io":     {TLD: "io", Channel: ChannelWhois, WhoisHost: "whois.nic.io"},
	"me":     {TLD: "me", Channel: ChannelWhois, WhoisHost: "whois.nic.me"},
	"co":     {TLD: "co", Channel: ChannelWhois, WhoisHost: "whois.registry.co"},
}

// SupportedTLDs returns the v1 TLD list, stable order.
func SupportedTLDs() []string {
	return []string{
		"com", "net", "org", "io", "me", "co",
		"dev", "app", "xyz", "info", "ai", "cc",
		"fun", "site", "online", "store", "tech", "pw", "cn",
	}
}

func NormalizeFQDN(fqdn string) (string, error) {
	s := strings.ToLower(strings.TrimSpace(fqdn))
	s = strings.TrimSuffix(s, ".")
	if s == "" || strings.Contains(s, " ") || strings.HasPrefix(s, ".") || strings.HasSuffix(s, ".") {
		return "", fmt.Errorf("invalid fqdn %q", fqdn)
	}
	if !strings.Contains(s, ".") {
		return "", fmt.Errorf("invalid fqdn %q", fqdn)
	}
	return s, nil
}

func tldOf(fqdn string) string {
	i := strings.LastIndex(fqdn, ".")
	if i < 0 || i == len(fqdn)-1 {
		return ""
	}
	return fqdn[i+1:]
}

func AdapterFor(fqdn string) (Adapter, error) {
	norm, err := NormalizeFQDN(fqdn)
	if err != nil {
		return Adapter{}, err
	}
	tld := tldOf(norm)
	a, ok := adapters[tld]
	if !ok {
		return Adapter{}, fmt.Errorf("unsupported tld %q", tld)
	}
	return a, nil
}
