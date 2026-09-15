# domain-probe

用 Go 写的 MCP 学习项目：域名查询工具，以及一份 `config.yaml`，用来接 OpenAI、Claude 或 OpenAI 兼容接口。

[English](README.md)

## 现状

阶段 0 只有域名内核（`internal/domain`）。把 WHOIS / RDAP 的响应分成 `available`、`registered`、`pending_delete`、`reserved`、`unknown`。MCP 服务和 LLM 壳还不在这个仓库里。

## v1 后缀

RDAP：`com` `net` `org` `dev` `app` `xyz` `info` `ai` `cc` `fun` `site` `online` `store` `tech` `pw`

WHOIS：`cn` `io` `me` `co`

其它后缀报错。打到错误注册局得到的 404 记为 `unknown`，不是 `available`。

## 开发

```bash
git clone https://github.com/hareai/domain-probe.git
cd domain-probe
go test ./...
```

需要 Go 1.24+。

## License

[MIT](LICENSE)
