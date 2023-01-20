#!/usr/bin/env bash
export GIN_MODE=release
go mod init gingonic-discord
go get github.com/gin-gonic/gin
go build -tags netgo -ldflags '-s -w' -o gingonic-discord