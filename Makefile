.DEFAULT_GOAL = help

vendor      := vendor
target      := target
web         := web
bin         := $(target)/bin
reports     := $(target)/reports

## build: Compile binaries.
.PHONY: build
build: build/backend build/frontend

.PHONY: build/backend
build/backend:
	@go build -o $(bin)/backend cmd/backend/main.go

.PHONY: build/frontend
build/frontend:
	@rm -f $(web)/app.wasm
	@npm install --prefix $(web)
	@GOOS=js GOARCH=wasm go build -o $(web)/app.wasm cmd/frontend/main.go
	@go build -o $(bin)/frontend cmd/frontend/main.go

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
