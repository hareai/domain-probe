package domain

import (
	"context"
	"io"
	"net/http"
	"os/exec"
	"time"
)

type Net struct {
	Client *http.Client
}

func (n Net) client() *http.Client {
	if n.Client != nil {
		return n.Client
	}
	return &http.Client{Timeout: 12 * time.Second}
}

func (n Net) RDAP(ctx context.Context, url string) (int, string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, "", err
	}
	req.Header.Set("User-Agent", "domain-probe/0.1")
	resp, err := n.client().Do(req)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
	return resp.StatusCode, string(body), nil
}

func (n Net) Whois(ctx context.Context, host, fqdn string) (string, error) {
	cmd := exec.CommandContext(ctx, "whois", "-h", host, fqdn)
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		return string(out), nil
	}
	return "", err
}
