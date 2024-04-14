.PHONY: default
default: setup

.PHONY: setup-buff
setup-buff:
	echo "setup-buff"
	GO111MODULE=on GOBIN=/usr/local/bin go install \
	github.com/bufbuild/buf/cmd/buf@v1.30.1

.PHONY: proto-generate
proto-generate:
	echo "proto-generate"
	buf generate

.PHONY: dependencies
dependencies:
	echo "dependencies"
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 ;
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 ;
	@export PATH="$$PATH:$(go env GOPATH)/bin" ;\
 	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
     github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 ;\
	go install github.com/vektra/mockery/v2@v2.42.2

.PHONY: build-migration
build-migration:
	echo "build-migration"
	go build -o bin/migration cmd/migration/main.go

.PHONY: build-api
build-api:
	echo "build-api"
	go build -o bin/api cmd/api/main.go

.PHONY: generate-mocks
generate-mocks:
    # Generate mocks in internal/mocks folder for all interfaces
	 mockery --dir ./internal/app/transaction --all --case snake --output internal/mocks/transaction/  --outpkg transaction_mocks
	 mockery --dir ./internal/app/cardholder --all --case snake --output internal/mocks/cardholder/  --outpkg cardholder_mocks

.PHONY: unit-tests
unit-tests:
	echo "unit-tests"
	go test -v -cover ./...

.PHONY: clean
clean:
	echo "clean"
	rm -rf ./internal/mocks/*
	rm -rf ./rpc/*
setup: dependencies setup-buff proto-generate