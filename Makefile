.PHONY: dep
dep: ## download and install dependencies
	@go get -u ./...
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	@rm api/tmdb.yml
	@curl https://raw.githubusercontent.com/ckatle/oas-tmdb/master/openapi.yml -o api/tmdb.yml

.PHONY: generate
generate: dep ## generate client code from openapi.yml
	@rm tmdb.gen.go
	@oapi-codegen -package tmdb api/tmdb.yml > tmdb.gen.go
	@go mod tidy