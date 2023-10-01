# syntax=docker/dockerfile:1

FROM golang:1.21.1-alpine3.18 as builder

RUN apk update && apk upgrade --no-cache

# Install Node.js and npm for tailwindcss
RUN apk add nodejs npm make curl brotli

WORKDIR /app

# Install dependencies. We want to cache this step as much as possible.
COPY go.mod go.sum package.json package-lock.json Makefile ./
RUN make init

COPY . .

# Build the application for production
RUN make assets
RUN make generate
RUN make build-prod

# https://github.com/GoogleContainerTools/distroless/blob/main/base/README.md
FROM gcr.io/distroless/static-debian12:nonroot-amd64

COPY --from=builder /app/bin/app /app/bin/app
COPY --from=builder /app/internal/assets/static /app/internal/assets/static

CMD ["/app/bin/app"]
