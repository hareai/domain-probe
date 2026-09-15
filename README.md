# domain-probe

A small Go project for learning the Model Context Protocol: domain lookup tools, and a `config.yaml` shell that talks to OpenAI, Claude, or an OpenAI-compatible endpoint.

[简体中文](README.zh-CN.md)

## Status

Stage 0 is the domain kernel only (`internal/domain`). It classifies a WHOIS or RDAP response into `available`, `registered`, `pending_delete`, `reserved`, or `unknown`. The MCP server and LLM shell are not in this tree yet.

## Supported TLDs (v1)

RDAP: `com` `net` `org` `dev` `app` `xyz` `info` `ai` `cc` `fun` `site` `online` `store` `tech` `pw`

WHOIS: `cn` `io` `me` `co`

Other suffixes return an error. A 404 from the wrong registry is `unknown`, not `available`.

## Develop

```bash
git clone https://github.com/hareai/domain-probe.git
cd domain-probe
go test ./...
```

Go 1.24+.

## License

[MIT](LICENSE)
