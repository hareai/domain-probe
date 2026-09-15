# domain-probe

用 Go 写的 MCP 学习项目：域名查询工具，以及一份 `config.yaml`，用来接 OpenAI、Claude 或 OpenAI 兼容接口。

[English](README.md)

## 现状

阶段 1 是 stdio MCP 服务，两个 tool。LLM 壳还不在这个仓库里。

## Tools

| 名字 | 输入 | 输出 |
|---|---|---|
| `lookup` | `name`（一个 FQDN） | `{name, status, channel}` |
| `lookup_batch` | `names`（列表）或 `label`（光名字，扫全部 v1 后缀） | `{results: [...]}` |

状态：`available`、`registered`、`pending_delete`、`reserved`、`unknown`。

## v1 后缀

RDAP：`com` `net` `org` `dev` `app` `xyz` `info` `ai` `cc` `fun` `site` `online` `store` `tech` `pw`

WHOIS：`cn` `io` `me` `co`

其它后缀报错。打到错误注册局得到的 404 记为 `unknown`，不是 `available`。批量查询间隔 2 秒。

## 运行

```bash
git clone https://github.com/hareai/domain-probe.git
cd domain-probe
go test ./...
go build -o domain-probe ./cmd/domain-probe
```

需要 Go 1.25+。

Claude Desktop / Cursor（stdio）：

```json
{
  "mcpServers": {
    "domain-probe": {
      "command": "/path/to/domain-probe"
    }
  }
}
```

Hermes（`config.yaml`）：

```yaml
mcp_servers:
  domain-probe:
    command: "/path/to/domain-probe"
```

## License

[MIT](LICENSE)
