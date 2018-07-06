#!/bin/bash

./scripts/listJSON.sh > list.json

go run cmd/generator/main.go $(pwd)/template/config.go.tmpl $(pwd)/config.go $(pwd)/list.json ~blacklist~

rm list.json
