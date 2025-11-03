# syntax=docker/dockerfile:1

## Stage 1: build the frontend assets
FROM node:20-alpine AS web-builder

WORKDIR /web

# Install dependencies using pnpm
COPY website/package.json website/pnpm-lock.yaml ./
RUN corepack enable \
    && pnpm install --no-frozen-lockfile

# Copy the remaining frontend sources and build
COPY website/ ./
RUN pnpm run build

## Stage 2: build the gateway binary
FROM golang:1.24 AS gateway-builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/gateway ./cmd/gateway

## Stage 3: assemble the runtime image
FROM gcr.io/distroless/base-debian12 AS gateway-runtime

WORKDIR /app

# Binary and static assets
COPY --from=gateway-builder /out/gateway /usr/local/bin/gateway
COPY --from=web-builder /web/dist ./website/dist
COPY configs/ ./configs/

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/gateway"]
