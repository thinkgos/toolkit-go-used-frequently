package main

import (
	// wire  依赖注入
	_ "github.com/google/wire/cmd/wire"
	// stringer 生成器
	_ "golang.org/x/tools/cmd/stringer"

	// buf toolkit
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking"
	_ "github.com/bufbuild/buf/cmd/protoc-gen-buf-lint"

	// proto toolkit
	_ "github.com/gogo/protobuf/protoc-gen-gogo"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

func main() {
	println("toolkit collection")
}
