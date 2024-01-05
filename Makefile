.PHONY: gen

gen: ## Generate mocks
	go generate -v ./...

test: ## Run tests
	go test -v ./...