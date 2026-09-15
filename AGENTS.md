# AI Project Contract

This file is binding for AI coding agents working in this repository.

## Before editing

- Read this file, README, CONTRIBUTING, and `git status`.
- Work on a topic branch. Do not commit to `main`.
- Do not guess product facts, credentials, or release intent.

## During editing

- Make the smallest coherent change.
- `internal/domain` is the domain kernel. It must not import MCP or LLM packages.
- Public repo: do not commit `.env`, live keys, or household/private domains (scan git history, not only HEAD).
- Keep tests, docs, and user-facing copy consistent in one PR.

## Before commit

- `go test ./...`
- Conventional Commit: `type(scope): imperative summary`
- Push a branch. Open a PR. Do not push `main` unless the owner authorizes an empty-repo first push or an emergency fix.

## Stop

Stop and ask when scope, destructive action, public behavior, or credentials are unclear.
