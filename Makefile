include .env

DEFAULT_GOAL: help

install: ;@ ## Install Setup
	@chmod +x scripts/install_dev.sh
	@./scripts/install_dev.sh
.PHONY: install

generate: install ;@ ## Generate Go Parser
	@echo "Generating Go parser with ANTLR..."
	@cd $(ANTLR_DIR) && \
	java -jar $(ANTLR_JAR) -Dlanguage=Go -o . -listener -visitor FSHLexer.g4 && \
	java -jar $(ANTLR_JAR) -Dlanguage=Go -o . -listener -visitor FSH.g4 && \
	mv *.go ../$(PARSER_DIR)
.PHONY: generate

fmt: ;@ ## Format Code
	@go fmt ./...
.PHONY: fmt

lint: ;@ ## Run Linter
	@@golangci-lint run ./...
.PHONY: lint

test: lint fmt ;@ ## Run Tests
	@go test ./... -v
.PHONY: test

build: generate ; @ ## Make Build
	@go build -o fshlint fshlint.go
.PHONY: build

help:
	@echo
	@echo
	@echo fshlint
	@echo
	@echo Commands
	@echo
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo
	@echo
.PHONY: help