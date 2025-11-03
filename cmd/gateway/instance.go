package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

const (
	apiPrefix = "/api/" // API prefix for all gRPC gateway endpoints
)

// Gateway wraps the grpc-gateway mux and provides HTTP endpoints for gRPC services
type Gateway struct {
	grpcServeMux *runtime.ServeMux
	grpcConn     *grpc.ClientConn
}

// NewGateway creates a new gateway instance
func NewGateway(
	grpcClientConn *grpc.ClientConn,
	grpcServeMux *runtime.ServeMux,
	grpcRegisterFuncs ...func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error,
) (*Gateway, error) {
	ctx := context.Background()
	for _, registerFunc := range grpcRegisterFuncs {
		if err := registerFunc(ctx, grpcServeMux, grpcClientConn); err != nil {
			_ = grpcClientConn.Close()
			return nil, fmt.Errorf("failed to register gRPC service: %w", err)
		}
	}

	return &Gateway{
		grpcServeMux: grpcServeMux,
		grpcConn:     grpcClientConn,
	}, nil
}

func (g *Gateway) APIMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle(apiPrefix, http.StripPrefix(strings.TrimSuffix(apiPrefix, "/"), g.grpcServeMux))
	return mux
}

func (g *Gateway) SPAMux(SPADistDir string) *http.ServeMux {
	mux := g.APIMux()
	if _, err := os.Stat(SPADistDir); err == nil {
		// Serve static files, with index.html as fallback for SPA
		fileServer := http.FileServer(http.Dir(SPADistDir))
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			path := filepath.Join(SPADistDir, r.URL.Path)

			// If file doesn't exist and it's not an API request
			// serve index.html for SPA routing
			if _, err := os.Stat(path); os.IsNotExist(err) {
				if !strings.HasPrefix(r.URL.Path, strings.TrimSuffix(apiPrefix, "/")) {
					http.ServeFile(w, r, filepath.Join(SPADistDir, "index.html"))
					return
				}
			}
			fileServer.ServeHTTP(w, r)
		})
		slog.Info("SPA serving enabled", "directory", SPADistDir)
	}
	return mux
}

// Close closes the gRPC connection
func (g *Gateway) Close() error {
	if g.grpcConn != nil {
		return g.grpcConn.Close()
	}
	return nil
}
