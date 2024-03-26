#!/bin/bash
go install github.com/cosmtrek/air@latest

# export GIN_MODE=release
air -c air.toml
