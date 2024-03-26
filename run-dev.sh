#!/bin/bash
go install github.com/cosmtrek/air@latest

# export GIN_MODE=release
export CGO_ENABLED=1 # Required for go-sqlite3
air -c air.toml
