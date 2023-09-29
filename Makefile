.PHONY: init assets generate run test build-prod

DIST_DIR := internal/assets/dist

# Install some important commands and tools
init:
	npm install
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/vektra/mockery/v2@latest
	go mod tidy

# Build the assets
assets:
	npx tailwindcss -i ./internal/assets/tailwind.css -o ./$(DIST_DIR)/styles.css
	curl -o $(DIST_DIR)/htmx.js https://unpkg.com/htmx.org@1.9.6/dist/htmx.min.js

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
	CGO_ENABLED=0 go build -o ./bin/app ./cmd/server/...

deploy:
	fly deploy