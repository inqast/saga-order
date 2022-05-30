GO_VERSION_SHORT:=$(shell echo `go version` | sed -E 's/.* go(.*) .*/\1/g')
ifneq ("1.17","$(shell printf "$(GO_VERSION_SHORT)\n1.17" | sort -V | head -1)")
$(error NEED GO VERSION >= 1.17. Found: $(GO_VERSION_SHORT))
endif

export GO111MODULE=on

SERVICE_PATH=inqast/saga-order

.PHONY: run
run:
	go run cmd/server/main.go

.PHONY: build-cart
build-cart:
	go mod download && CGO_ENABLED=0  go build \
		-tags='no_mysql no_sqlite3' \
		-ldflags=" \
			-X 'github.com/$(SERVICE_PATH)/internal/config.version=$(VERSION)' \
			-X 'github.com/$(SERVICE_PATH)/internal/config.commitHash=$(COMMIT_HASH)' \
		" \
		-o ./bin/service$(shell go env GOEXE) ./cmd/cart/main.go

protobuf-cart:
	protoc --proto_path=./api/cart \
		--go_out=pkg/api/cart \
		--go_opt=paths=source_relative \
		--go-grpc_out=pkg/api/cart \
		--go-grpc_opt=paths=source_relative \
		./api/cart/api.proto
