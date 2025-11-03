#!/bin/bash

./scripts/install_brew.sh

brew install bufbuild/buf/buf

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

buf dep update

nvm use
if ! command -v pnpm >/dev/null 2>&1; then
    npm install -g pnpm
fi
pnpm install
(cd website && pnpm install)

./scripts/download_test_data.sh