BIN := ddd

all: fmt lint build test

.PHONY: fmt
fmt: ## format golang source code
	find . -name "*.go" -not -path "./vendor/*" -not -path "./git/*" -not -path "*/mock/*" | xargs goimports -w
	find . -name "*.go" -not -path "./vendor/*" -not -path "./git/*" -not -path "*/mock/*" | xargs gofmt -w

.PHONY: build
build: ## build code
	go build -o $(BIN) main.go

.PHONY: test
test: ## test for project
	go test ./...

.PHONY: run
run: ## run code
	go run main.go

.PHONY: lint
lint: ## golang lint code
	golangci-lint  run -j 10 --disable-all \
    --enable=gofmt \
    --enable=golint \
    --enable=gosimple \
    --enable=ineffassign \
    --enable=vet \
    --enable=misspell \
    --enable=unconvert \
    --exclude='should have comment' \
    --exclude='and that stutters;' \
     . 2>&1 | grep -v 'ALL_CAPS\|OP_' 2>&1 | tee /dev/stderr

.PHONY: strict_lint
strict_lint: ## golang lint code, do not ignore any warning
	golangci-lint run ./...

.PHONY: swagger
swagger: ## Generate swagger files into ./api. install swag with `go get -u github.com/swaggo/swag/cmd/swag`
	swag init

.PHONY: help
help:
	@ grep -h -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean:
	rm -rf $(BIN)