export

LOCAL_BIN:=$(CURDIR)/bin
PATH:=$(LOCAL_BIN):$(PATH)

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

setup: ## Ensure the go.mod file is clean and updated with the project dependencies.
	pip3 install pre-commit commitizen
	go mod tidy
.PHONY: setup

install-rie: ## install AWS Lambda RIE at local folder
	chmod +x install-rie.sh
	./install-rie.sh
.PHONY: install-rie

compose-up: ## Build and start service using docker-compose.
	docker-compose up --build -d
.PHONY: compose-up

compose-down: ## Stop and remove all services defined in docker-compose.
	docker-compose down --remove-orphans
.PHONY: compose-down

compose-logs: ## Show its logs using docker-compose.
	docker-compose logs --tail=30 -f
.PHONY: compose-logs

fmt: ## Format and vet Go code using gofmt, goimports and go vet.
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Running go vet..."
	@go vet ./...
	@if command -v goimports > /dev/null; then \
		goimports -w -local github.com/javiertelioz/aws-lambda-golang .; \
	else \
		echo "goimports not found, skipping..."; \
	fi
.PHONY: fmt

linter: ## Run the golangci-lint on the project source code to detect style issues or errors.
	golangci-lint run
.PHONY: linter

test: ## Run all tests with race detection, parallel execution and shuffled order.
	go clean -testcache
	go test -v -race -parallel 4 -shuffle=on ./test/...
.PHONY: test

coverage: ## Generate and visualize a test coverage report in HTML format.
	@mkdir -p coverage
	go clean -testcache
	go test -v -race -cover -covermode=atomic ./test/... -coverpkg=./pkg/... -coverprofile=coverage/coverage.out -shuffle=on
	go tool cover -func=coverage/coverage.out
	go tool cover -html=coverage/coverage.out -o coverage/coverage.html
.PHONY: coverage

pre-commit: ## Select files to add to staging area using fzf (interactive) or add all changes.
	@echo "📦 Selecting files to add..."
	@if command -v fzf > /dev/null; then \
		git status --short | \
		fzf --multi \
		    --height=60% \
		    --border=rounded \
		    --prompt="Select files > " \
		    --header="TAB: select/deselect | ENTER: confirm | ESC: cancel" \
		    --preview="echo {} | awk '{print \$$2}' | xargs git diff --color=always" \
		    --preview-window=right:60%:wrap | \
		awk '{print $$2}' | \
		xargs -I {} git add "{}"; \
		if [ $$? -eq 0 ]; then \
			echo "✅ Files added to staging area"; \
			echo "Staged files:"; \
			git diff --name-only --cached; \
		else \
			echo "❌ No files selected or operation cancelled"; \
			exit 1; \
		fi; \
	else \
		echo "⚠️  fzf not found. Adding all changes..."; \
		git add .; \
	fi
	@echo "📝 Creating commit with commitizen..."
	@cz commit
.PHONY: pre-commit

commit: ## Run all quality checks (fmt, vet, linter, tests, coverage) and create a commit with commitizen.
	@echo "🔍 Running quality checks before commit..."
	@echo "\n📝 Step 1/5: Formatting and vetting code..."
	@$(MAKE) fmt
	@echo "\n✅ Step 2/5: Running linter..."
	@$(MAKE) linter
	@echo "\n🧪 Step 3/5: Running tests..."
	@$(MAKE) test
	@echo "\n📊 Step 4/5: Generating coverage..."
	@$(MAKE) coverage
	@echo "\n✨ All quality checks passed!"
	@echo "\n📦 Step 5/5: Adding files to commit..."
	@$(MAKE) pre-commit
.PHONY: commit
