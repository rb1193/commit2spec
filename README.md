# Commit2Spec

Commit2Spec is a CLI utility that accepts a git commit hash, or a range of Git commit hashes, as arguments, and generates a spec file in markdown for each commit which could be passed to a coding assistant such as Claude Code to implement.

Commit2Spec uses Claude Opus to generate the spec from the contents of the commit. It reads an API key for Claude from an environment variable.

Commit2Spec is written in go and is a binary executable and uses the Cobra library to parse the CLI command.

Git commits are read using the go-git library rather than shelling out to git.


