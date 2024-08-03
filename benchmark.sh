#!/usr/bin/env sh

go test ./... -json -run=Benchmark -bench=. -benchmem | tee benchmark.jsonl