.PHONY: *

clean:
	@go clean -cache
	@go clean -modcache

dep: ## download and install dependencies
	@echo "download dependencies"
	@go get -u ./...
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

generate: dep ## generate client code from openapi.yml
	@rm tmdb.gen.go
	@oapi-codegen -package tmdb api/openapi.yml > tmdb.gen.go
	@go mod tidy

pc: pca pcr

pca: ## Updating hooks automatically
	@pre-commit autoupdate

pcr: ## Run against all the files
	@pre-commit run -a