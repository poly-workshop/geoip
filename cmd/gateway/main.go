package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	geoip_v1_pb "github.com/poly-workshop/geoip/gen/go/geoip/v1"
	"github.com/poly-workshop/geoip/internal/configs"
	"github.com/poly-workshop/go-webmods/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	cmdName = "gateway"
)

func init() {
	app.Init(cmdName)
}

func main() {
	cfg := configs.LoadGatewayConfig()

	// Create gRPC connection
	conn, err := grpc.NewClient(
		cfg.GRPCEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	slog.Info("Setup gRPC connection", "endpoint", cfg.GRPCEndpoint)

	// Create gateway mux with custom options
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
			switch key {
			case "Authorization":
				return key, true
			default:
				return "", false
			}
		}),
	)

	// Create gateway instance
	gateway, err := NewGateway(
		conn, mux,
		geoip_v1_pb.RegisterGeoIPServiceHandler,
	)
	if err != nil {
		log.Fatalf("failed to create gateway: %v", err)
	}
	defer func() { _ = gateway.Close() }()

	// Create HTTP server
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: gateway.SPAMux("website/dist"),
	}
	slog.Info("HTTP gateway server started", "port", cfg.Port)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
