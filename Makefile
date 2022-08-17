.DEFAULT_GOAL = help

vendor      := vendor
target      := target
web         := web
bin         := $(target)/bin
reports     := $(target)/reports

## generate: Run generators.
.PHONY: generate
generate: go/generate

.PHONY: go/generate
go/generate:
	@go generate ./...

## build: Compile binaries.
.PHONY: build
build: build/server build/ui

.PHONY: build/server
build/server:
	@go build -o $(bin)/server cmd/server/main.go

.PHONY: build/ui
build/ui:
	@rm -f $(web)/app.wasm
	@npm install --silent --prefix $(web)
	@GOOS=js GOARCH=wasm go build -o $(web)/app.wasm cmd/ui/main.go
	@go build -o $(bin)/ui cmd/ui/main.go

$(reports):
	@mkdir -p $@

## lint: Run static analysis.
.PHONY: lint
lint: go/lint

.PHONY: go/lint
go/lint:
	@golangci-lint run

## tests: Run tests.
.PHONY: test tests
test tests: go/test

.PHONY: go/test
go/test: $(go_src) | $(reports)
	@go test -v -covermode=atomic -coverprofile=$(reports)/cover.out ./...

## clean: Remove created resources.
.PHONY: clean
clean:
	@rm -rf $(vendor) $(target)

## help: Display available targets.
.PHONY: help
help: Makefile
	@echo "Usage: make [target]"
	@echo
	@echo "Targets:"
	@sed -n 's/^## //p' $< | awk -F ':' '{printf "  %-20s%s\n",$$1,$$2}'
