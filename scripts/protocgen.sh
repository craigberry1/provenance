#!/usr/bin/env bash

#== Requirements ==
#
## make sure your `go env GOPATH` is in the `$PATH`
## Install:
## + latest buf (v1.0.0-rc11 or later)
## + protobuf v3
#
## All protoc dependencies must be installed not in the module scope
## currently we must use grpc-gateway v1
# cd ~
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.16.0
# go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest

set -eo pipefail

repo_root="$( cd "$( dirname "${BASH_SOURCE:-$0}" )/.."; pwd -P )"
cd "$repo_root/proto"

# Generate Go code from the proto files.
buf generate --template buf.gen.gogo.yaml

# Generate docs using protoc-gen-doc
buf generate --template buf.gen.docs.yaml

cd "$repo_root"

# move proto files to the right places
cp -r github.com/provenance-io/provenance/* ./
rm -rf github.com
