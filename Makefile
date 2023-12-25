# ファイルが分割されているとき
# bundle:
# 	mkdir -p ./build
# 	docker run --rm -v $$PWD:/spec redocly/cli bundle openapi.yaml > ./build/openapi.yaml

install-backend-types-generator:
	@which openapi-codegen > /dev/null || go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

generate-backend-types:
	mkdir -p ./gen-pkg
	oapi-codegen -generate types -package types build/openapi.yaml > ./gen-pkg/types.gen.go

generate-backend-all:
	mkdir -p ./gen-pkg
	oapi-codegen -package all build/openapi.yaml > ./gen-pkg/all.gen.go

generate-backend-types-and-client:
	mkdir -p ./gen-pkg
	oapi-codegen -generate types,client -package typesclient build/openapi.yaml > ./gen-pkg/typesclient.gen.go

generate-backend-server:
	mkdir -p ./gen-pkg
	oapi-codegen -generate server -package server build/openapi.yaml > ./gen-pkg/server.gen.go
