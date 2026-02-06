# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Commit2Spec is a Go CLI utility that generates markdown spec files from git commits. It takes commit hashes (single or range) as arguments and uses Claude Opus to generate implementation specs that can be passed to coding assistants.

## Development Environment

Uses Mise for tool version management. Run `mise install` to set up Go (latest version).

## Build Commands

```bash
go build -o commit2spec .    # Build binary
go run .                      # Run without building
go test ./...                 # Run all tests
go test -v ./... -run TestName  # Run specific test
```

## Configuration

The CLI reads the Claude API key from an environment variable (implementation will define which one).

## Architecture Notes

This project is in early development. When implementing:
- Main entry point should be in `main.go` or `cmd/` directory
- Consider a `pkg/` or `internal/` directory for reusable packages
- Separate concerns: git operations, Claude API client, spec generation, CLI parsing
