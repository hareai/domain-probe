# domain-probe

A small Go project for learning the Model Context Protocol: domain lookup tools, and a `config.yaml` shell that talks to OpenAI, Claude, or an OpenAI-compatible endpoint.

[简体中文](README.zh-CN.md)

## Status

Stage 1 is a stdio MCP server with two tools. The LLM shell is not in this tree yet.

## Tools

| Name | Input | Output |
|---|---|---|
| `lookup` | `name` (one FQDN) | `{name, status, channel}` |
| `lookup_batch` | `names` (list) or `label` (bare word, every v1 suffix) | `{results: [...]}` |

Status is one of `available`, `registered`, `pending_delete`, `reserved`, `unknown`.

## Supported TLDs (v1)

RDAP: `com` `net` `org` `dev` `app` `xyz` `info` `ai` `cc` `fun` `site` `online` `store` `tech` `pw`

WHOIS: `cn` `io` `me` `co`

Other suffixes return an error. A 404 from the wrong registry is `unknown`, not `available`. Batch lookups wait 2 seconds between names.

## Run

```bash
git clone https://github.com/hareai/domain-probe.git
cd domain-probe
go test ./...
go build -o domain-probe ./cmd/domain-probe
```

Go 1.25+.

Claude Desktop / Cursor (stdio):

```json
{
  "mcpServers": {
    "domain-probe": {
      "command": "/path/to/domain-probe"
    }
  }
}
```

Hermes (`config.yaml`):

```yaml
mcp_servers:
  domain-probe:
    command: "/path/to/domain-probe"
```

## License

[MIT](LICENSE)
