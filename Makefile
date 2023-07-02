.PHONY: gen
gen: ## generate client code from openapi.yml
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	@oapi-codegen -package tmdb https://raw.githubusercontent.com/ckatle/oas-tmdb/master/openapi.yml > tmdb.gen.go
	@go get -u ./...
	@go mod tidy