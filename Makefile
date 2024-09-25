.PHONY: all
all: generate tidy fmt build

.PHONY: build
build:
	go build ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: generate
generate:
	bash hack/codegen.sh
