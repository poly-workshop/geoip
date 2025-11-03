# syntax=docker/dockerfile:1

## Build stage: compile the gRPC server binary
FROM golang:1.25 AS grpc-builder

WORKDIR /src

# Preload go modules
COPY go.mod go.sum ./
RUN go mod download

# Install tooling required for auxiliary scripts
RUN apt-get update \
	&& apt-get install -y --no-install-recommends wget ca-certificates \
	&& rm -rf /var/lib/apt/lists/*

# Copy the application source
COPY . .

# Fetch open-source GeoIP datasets for local builds
RUN bash ./scripts/download_opensource_data.sh

# Build the gRPC server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/grpc-server ./cmd/grpc-server

## Runtime stage
FROM gcr.io/distroless/base-debian12 AS grpc-runtime

WORKDIR /app

# Copy binary and default configs
COPY --from=grpc-builder /out/grpc-server /usr/local/bin/grpc-server
COPY configs/ ./configs/
COPY --from=grpc-builder /src/data ./data/

EXPOSE 50051

ENTRYPOINT ["/usr/local/bin/grpc-server"]
