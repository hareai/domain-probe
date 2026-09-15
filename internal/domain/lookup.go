package domain

import (
	"context"
	"strings"
	"time"
)

type Record struct {
	Name    string  `json:"name"`
	Status  Status  `json:"status"`
	Channel Channel `json:"channel"`
}

type Getter interface {
	RDAP(ctx context.Context, url string) (status int, body string, err error)
	Whois(ctx context.Context, host, fqdn string) (text string, err error)
}

type Probe struct {
	Get Getter
	Gap time.Duration
}

func (p Probe) Lookup(ctx context.Context, name string) (Record, error) {
	a, err := AdapterFor(name)
	if err != nil {
		return Record{}, err
	}
	norm, _ := NormalizeFQDN(name)
	rec := Record{Name: norm, Channel: a.Channel}

	if a.Channel == ChannelRDAP {
		code, _, err := p.Get.RDAP(ctx, a.RDAPURL+norm)
		if err != nil {
			rec.Status = StatusUnknown
			return rec, nil
		}
		rec.Status = ClassifyRDAP(code)
		return rec, nil
	}

	text, err := p.Get.Whois(ctx, a.WhoisHost, norm)
	if err != nil {
		rec.Status = StatusUnknown
		return rec, nil
	}
	rec.Status = ClassifyWhois(a.TLD, norm, text)
	return rec, nil
}

func (p Probe) LookupMany(ctx context.Context, names []string) []Record {
	out := make([]Record, 0, len(names))
	for i, name := range names {
		if i > 0 && p.Gap > 0 {
			t := time.NewTimer(p.Gap)
			select {
			case <-ctx.Done():
				t.Stop()
				return out
			case <-t.C:
			}
		}
		rec, err := p.Lookup(ctx, name)
		if err != nil {
			out = append(out, Record{Name: name, Status: StatusUnknown})
			continue
		}
		out = append(out, rec)
	}
	return out
}

func Expand(label string) []string {
	label = strings.ToLower(strings.TrimSpace(label))
	tlds := SupportedTLDs()
	names := make([]string, 0, len(tlds))
	for _, tld := range tlds {
		names = append(names, label+"."+tld)
	}
	return names
}
