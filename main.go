package main

import (
	// wire  依赖注入
	_ "github.com/google/wire/cmd/wire"
	// stringer 生成器
	_ "golang.org/x/tools/cmd/stringer"

	// buf 相关
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking"
	_ "github.com/bufbuild/buf/cmd/protoc-gen-buf-lint"

	// proto 相关
	_ "github.com/gogo/protobuf/protoc-gen-gogo"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"

	// 数据库相关
	// _ "github.com/github/gh-ost/go/cmd/gh-ost" // online DDL
	_ "github.com/xo/usql" // 通用sql接口

	// uml
	_ "github.com/jfeliu007/goplantuml/cmd/goplantuml"
)

func main() {
	println("toolkit collection")
}
