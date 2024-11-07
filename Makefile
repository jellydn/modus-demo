.PHONY: install-cli
install-cli:
	@echo "Installing modus..."
	curl install.hypermode.com/modus.sh -sSfL | bash

.PHONY: dev
dev:
	@echo "Running dev server..."
	@modus dev

.PHONY: test
test:
	@echo "Running tests..."
	@go test

.PHONY: help
help:
	@echo "Available commands:"
	@echo "  install-cli - Install modus CLI"
	@echo "  dev         - Run dev server"
	@echo "  test        - Run tests"
	@echo "  help        - Show this help message"
