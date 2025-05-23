.PHONY: help tidy build sqlc install-tools

# Default target
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## ' Makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

tidy: ## Tidy the project
build: ## Build the project
serve: ## Serve the project
sqlc: ## Generate sqlc types and queries

tidy:
	@echo "Tidying the project..."
	@go mod tidy

build:
	@make tidy
	@echo "Building the project..."
	@go build -o bin/api cmd/api/main.go

serve:
	@echo "Serving the project..."
	@bin/api

sqlc:
	@echo "Generating sqlc types and queries..."
	@cd sqlc generate

# Tool installation
install-tools: ## Install all required development tools
	@echo "Installing required tools..."
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@go install github.com/mikefarah/yq/v4@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "Tools installed successfully!"
