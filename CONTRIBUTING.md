# Contributing

Thanks for contributing. This project uses a focused branch → PR → CI → merge workflow, including for solo development.

## Workflow

1. Start from an up-to-date `main`.
2. Create `feat/*`, `fix/*`, `refactor/*`, `docs/*`, `test/*`, or `ci/*`.
3. Keep one logical task per branch and PR.
4. Run `go test ./...` before pushing.
5. Open a PR.
6. Merge only after required CI checks are green.
7. Prefer squash merge, then delete the branch.

## Commits

Use Conventional Commits:

```text
type(scope): imperative summary
```

Allowed types: `feat`, `fix`, `refactor`, `docs`, `test`, `ci`, `build`, `chore`, `perf`, `revert`.

## Required with behavior changes

- Tests added or updated.
- README/docs updated when behavior or usage changes.
- `CHANGELOG.md` updated under `[Unreleased]` for user-facing changes.
- No secrets or private infrastructure in the diff.

## Security

Do not report vulnerabilities in public issues. Follow `SECURITY.md`.
