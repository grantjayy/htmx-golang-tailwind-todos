.PHONY: init assets generate run test build-prod

STATIC_DIR := internal/assets/static

# Install some important commands and tools
init:
	npm install
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/vektra/mockery/v2@latest
	go mod tidy

# Build the assets
assets:
	npx tailwindcss -i ./internal/assets/tailwind.css -o ./$(STATIC_DIR)/styles.css
	curl -o $(STATIC_DIR)/htmx.js https://unpkg.com/htmx.org@1.9.6/dist/htmx.min.js
	curl -o $(STATIC_DIR)/hyperscript.js https://unpkg.com/hyperscript.org@0.9.11/dist/_hyperscript.min.js
	find ./internal/assets/static/ -type f \( -name '*.js' -o -name '*.css' \) -exec brotli --best --force {} \; -exec mv "{}.br" "{}" \;

# Generate stuff and things
generate:
	go generate ./...

# Run the application
run: assets generate
	go run ./cmd/server/...

# Run the tests
test:
	go test ./...

# Build production binary
build-prod:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/app ./cmd/server/...

deploy-prod:
	fly deploy --config fly.prod.toml --remote-only --strategy immediate