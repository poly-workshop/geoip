package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"

	"github.com/go-redis/redis_rate/v10"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	geoip_v1_pb "github.com/poly-workshop/geoip/gen/go/geoip/v1"
	"github.com/poly-workshop/geoip/internal/configs"
	geoip_v1 "github.com/poly-workshop/geoip/internal/geoip-v1"
	"github.com/poly-workshop/geoip/internal/maxmind"
	"github.com/poly-workshop/go-webmods/app"
	grpc_utils "github.com/poly-workshop/go-webmods/grpc-utils"
	redis_client "github.com/poly-workshop/go-webmods/redis-client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	cmdName = "grpc-server"
)

func init() {
	app.Init(cmdName)
}

// InterceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(
		func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
			l.Log(ctx, slog.Level(lvl), msg, fields...)
		},
	)
}

func main() {
	cfg := configs.LoadGRPCServerConfig()

	if err := maxmind.EnsureDatabases(context.Background(), cfg); err != nil {
		log.Fatalf("failed to ensure MaxMind databases: %v", err)
	}

	var (
		redisLimiter *redis_rate.Limiter
		redisLimit   redis_rate.Limit
	)

	if len(cfg.RedisURLs) > 0 && cfg.RateLimitPerMinute > 0 {
		rdb := redis_client.NewRDB(redis_client.Config{
			Urls:     cfg.RedisURLs,
			Password: cfg.RedisPassword,
		})
		if err := rdb.Ping(context.Background()).Err(); err != nil {
			log.Fatalf("failed to connect to redis: %v", err)
		}

		redisLimiter = redis_rate.NewLimiter(rdb)
		redisLimit = redis_rate.PerMinute(cfg.RateLimitPerMinute)
		if cfg.RateLimitBurst > 0 {
			redisLimit.Burst = cfg.RateLimitBurst
		}

		slog.Info(
			"rate limiting enabled",
			"per_minute", redisLimit.Rate,
			"burst", redisLimit.Burst,
		)
	} else {
		slog.Info("rate limiting disabled")
	}

	geoipService := geoip_v1.NewService(cfg, redisLimiter, redisLimit)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_utils.BuildRequestIDInterceptor(),
			logging.UnaryServerInterceptor(InterceptorLogger(slog.Default())),
		),
	)
	geoip_v1_pb.RegisterGeoIPServiceServer(grpcServer, geoipService)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen on gRPC port: %v", err)
	}

	slog.Info("gRPC server started", "port", cfg.Port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}
