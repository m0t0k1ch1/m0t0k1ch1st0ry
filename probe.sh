#!/bin/sh
# temporary probe (to be reverted)
exec > static/probe.txt 2>&1
set -x
C="$PWD/node_modules/.cache"
mkdir -p "$C"
ls "$C"; du -sh "$C"/* 2>/dev/null
date +%s
if [ ! -x "$C/go/bin/go" ]; then curl -sSL https://go.dev/dl/go1.27.4.linux-amd64.tar.gz | tar -xz -C "$C"; fi
date +%s
export PATH="$C/go/bin:$PATH" GOCACHE="$C/go-build" GOMODCACHE="$C/go-mod" GOTOOLCHAIN=local
go version
go tool hugo version
date +%s
du -sh "$C"/*
