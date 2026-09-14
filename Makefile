# RealWorld (Conduit) — Go + Vue learning template.
# Every workflow in this repo goes through a make target. Run `make` for the list.

SHELL := /bin/bash
.DEFAULT_GOAL := help

# Load .env if present so DATABASE_URL etc. are available to every target.
ifneq (,$(wildcard .env))
include .env
export
endif

DATABASE_URL ?= postgres://conduit:conduit@localhost:5432/conduit?sslmode=disable
PORT         ?= 8080
API_DIR      := api
WEB_DIR      := web
SPECS_DIR    := .specs

export DATABASE_URL
export PORT

.PHONY: help
help: ## Show this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage: make <target>\n\n"} \
		/^[a-zA-Z0-9_.-]+:.*?##/ { printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2 } \
		/^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) }' $(MAKEFILE_LIST)
	@echo ""

##@ Setup

.PHONY: setup
setup: .env tools specs-sync deps ## One-time setup: env file, CLI tools, upstream specs, dependencies
	@echo ""
	@echo "Setup complete. Next:  make up && make migrate-up && make dev"

.env:
	@cp .env.example .env && echo "Created .env from .env.example"

.PHONY: tools
tools: ## Install the Go CLI tools this repo uses (goose, sqlc, golangci-lint, air)
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
	go install github.com/air-verse/air@latest
	@echo ""
	@echo "Also install Hurl (API conformance runner): https://hurl.dev/docs/installation.html"

.PHONY: deps
deps: ## Download Go modules and install web + e2e dependencies
	cd $(API_DIR) && go mod download
	npm install
	cd $(WEB_DIR) && npm install
	npx playwright install chromium

.PHONY: specs-sync
specs-sync: ## Fetch the upstream RealWorld conformance suites into .specs/
	./scripts/sync-specs.sh

##@ Infrastructure

.PHONY: up
up: ## Start Postgres (and Adminer on :8081), waiting until healthy
	docker compose up -d --wait

.PHONY: down
down: ## Stop containers
	docker compose down

.PHONY: reset-db
reset-db: ## Destroy the database volume and re-apply all migrations
	docker compose down -v
	$(MAKE) up
	$(MAKE) migrate-up

.PHONY: psql
psql: ## Open a psql shell against the dev database
	docker compose exec postgres psql -U $${POSTGRES_USER:-conduit} -d $${POSTGRES_DB:-conduit}

##@ Database

GOOSE = goose -dir $(API_DIR)/migrations postgres "$(DATABASE_URL)"

.PHONY: migrate-up
migrate-up: ## Apply all pending migrations
	$(GOOSE) up

.PHONY: migrate-down
migrate-down: ## Roll back the most recent migration
	$(GOOSE) down

.PHONY: migrate-status
migrate-status: ## Show migration status
	$(GOOSE) status

.PHONY: migrate-new
migrate-new: ## Create a migration: make migrate-new name=add_articles
	@test -n "$(name)" || (echo "usage: make migrate-new name=<snake_case_name>" && exit 1)
	goose -dir $(API_DIR)/migrations create $(name) sql

.PHONY: sqlc
sqlc: ## Regenerate internal/store from migrations + queries
	cd $(API_DIR) && sqlc generate

##@ Run

.PHONY: api
api: ## Run the Go API (PORT env, default 8080)
	cd $(API_DIR) && go run ./cmd/api

.PHONY: api-watch
api-watch: ## Run the Go API with live reload (requires air)
	cd $(API_DIR) && air

.PHONY: web
web: ## Run the Vite dev server on :5173
	cd $(WEB_DIR) && npm run dev

.PHONY: dev
dev: ## Run API and web together (Ctrl-C stops both)
	@echo "API  -> http://localhost:$(PORT)"
	@echo "Web  -> http://localhost:5173"
	@trap 'kill 0' EXIT INT TERM; \
		( cd $(API_DIR) && go run ./cmd/api ) & \
		( cd $(WEB_DIR) && npm run dev ) & \
		wait

##@ Quality

.PHONY: lint
lint: lint-go lint-web ## Lint everything

.PHONY: lint-go
lint-go: ## Vet + golangci-lint the Go module
	cd $(API_DIR) && go vet ./... && golangci-lint run

.PHONY: lint-web
lint-web: ## ESLint + vue-tsc type check
	cd $(WEB_DIR) && npm run lint && npm run typecheck

.PHONY: fmt
fmt: ## Format Go and web sources
	cd $(API_DIR) && go fmt ./...
	cd $(WEB_DIR) && npm run format

.PHONY: test
test: test-go test-web ## Run unit tests

.PHONY: test-go
test-go: ## Go unit tests with race detector
	cd $(API_DIR) && go test -race ./...

.PHONY: test-web
test-web: ## Vitest unit tests
	cd $(WEB_DIR) && npm run test:unit

.PHONY: build
build: ## Build both halves
	cd $(API_DIR) && go build -o bin/api ./cmd/api
	cd $(WEB_DIR) && npm run build

##@ Conformance (your to-do list)

.PHONY: verify
verify: verify-api verify-web ## Run both conformance suites

.PHONY: verify-api
verify-api: $(SPECS_DIR) ## Hurl API suite against a running API — the backend to-do list
	HOST=http://localhost:$(PORT)/api $(SPECS_DIR)/api/run-api-tests-hurl.sh

.PHONY: verify-web
verify-web: $(SPECS_DIR) ## Playwright e2e suite (starts API + web itself) — the frontend to-do list
	TEST_MODE=fullstack npx playwright test

.PHONY: verify-report
verify-report: ## Open the last Playwright HTML report
	npx playwright show-report

$(SPECS_DIR):
	@echo "Upstream specs missing. Run: make specs-sync" && exit 1

.PHONY: clean
clean: ## Remove build output and test artifacts
	rm -rf $(API_DIR)/bin $(API_DIR)/tmp $(WEB_DIR)/dist playwright-report test-results
