#!/usr/bin/env bash
pushd / 1>/dev/null ; find . -name print.go -type f 2>/dev/null | pipepeek 1>/dev/null ; popd 1>/dev/null
