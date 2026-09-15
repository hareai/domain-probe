# Changelog

All notable changes are documented here.

Format based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), versioned by [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Domain kernel (`internal/domain`): TLD adapters and WHOIS/RDAP classifiers for the v1 suffix list.
- stdio MCP server (`cmd/domain-probe`) with `lookup` and `lookup_batch`.
