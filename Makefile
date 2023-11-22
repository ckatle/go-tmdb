.PHONY: *

clean:
	@go clean -cache
	@go clean -modcache

generate: ## generate client code from openapi.yml
	@go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
	@rm ./pkg/tmdb/*
	@oapi-codegen --config=./configs/api/models.yml ./api/openapi.yml
	@oapi-codegen --config=./configs/api/client.yml ./api/openapi.yml
	#@oapi-codegen -package tmdb api/openapi.yml > ./pkg/tmdb/client.gen.go
	@go mod tidy

pc: pca pcr

pca: ## Updating hooks automatically
	@pre-commit autoupdate

pcr: ## Run against all the files
	@pre-commit run -a