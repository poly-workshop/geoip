package geoip_v1

import (
	"context"
	"log/slog"
	"net/netip"
	"os"

	"github.com/go-redis/redis_rate/v10"
	geoip2 "github.com/oschwald/geoip2-golang/v2"
	geoip_v1_pb "github.com/poly-workshop/geoip/gen/go/geoip/v1"
	"github.com/poly-workshop/geoip/internal/configs"
	"github.com/poly-workshop/geoip/internal/geoip-v1/convert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	datasetCity           = "city"
	datasetCountry        = "country"
	datasetEnterprise     = "enterprise"
	datasetAnonymousIP    = "anonymous_ip"
	datasetASN            = "asn"
	datasetConnectionType = "connection_type"
	datasetDomain         = "domain"
	datasetISP            = "isp"
)

type datasetConfig struct {
	path      string
	available bool
}

type Service struct {
	geoip_v1_pb.UnimplementedGeoIPServiceServer
	datasets  map[string]datasetConfig
	limiter   *redis_rate.Limiter
	rateLimit redis_rate.Limit
}

func NewService(
	cfg configs.GRPCServerConfig,
	limiter *redis_rate.Limiter,
	limit redis_rate.Limit,
) *Service {
	datasets := map[string]datasetConfig{
		datasetCity:           buildDatasetConfig(cfg.CityDBPath),
		datasetCountry:        buildDatasetConfig(cfg.CountryDBPath),
		datasetEnterprise:     buildDatasetConfig(cfg.EnterpriseDBPath),
		datasetAnonymousIP:    buildDatasetConfig(cfg.AnonymousIPDBPath),
		datasetASN:            buildDatasetConfig(cfg.ASNDBPath),
		datasetConnectionType: buildDatasetConfig(cfg.ConnectionTypeDBPath),
		datasetDomain:         buildDatasetConfig(cfg.DomainDBPath),
		datasetISP:            buildDatasetConfig(cfg.ISPDBPath),
	}

	service := &Service{datasets: datasets}
	if limiter != nil && !limit.IsZero() {
		service.limiter = limiter
		service.rateLimit = limit
	}

	return service
}

func buildDatasetConfig(path string) datasetConfig {
	if path == "" {
		return datasetConfig{}
	}

	return datasetConfig{
		path:      path,
		available: fileExists(path),
	}
}

func (s *Service) GetCity(
	ctx context.Context,
	req *geoip_v1_pb.GetCityRequest,
) (*geoip_v1_pb.GetCityResponse, error) {
	if err := s.enforceRateLimit(ctx, "GetCity"); err != nil {
		return nil, err
	}

	path, err := s.datasetPath(datasetCity)
	if err != nil {
		return nil, err
	}

	ip, err := parseIP(req.GetIpAddress())
	if err != nil {
		return nil, err
	}

	record, err := withReader(path, func(r *geoip2.Reader) (*geoip2.City, error) {
		return r.City(ip)
	})
	if err != nil {
		return nil, err
	}

	return &geoip_v1_pb.GetCityResponse{City: convert.City(record)}, nil
}

func (s *Service) GetCountry(
	ctx context.Context,
	req *geoip_v1_pb.GetCountryRequest,
) (*geoip_v1_pb.GetCountryResponse, error) {
	if err := s.enforceRateLimit(ctx, "GetCountry"); err != nil {
		return nil, err
	}

	path, err := s.datasetPath(datasetCountry)
	if err != nil {
		return nil, err
	}

	ip, err := parseIP(req.GetIpAddress())
	if err != nil {
		return nil, err
	}

	record, err := withReader(path, func(r *geoip2.Reader) (*geoip2.Country, error) {
		return r.Country(ip)
	})
	if err != nil {
		return nil, err
	}

	return &geoip_v1_pb.GetCountryResponse{Country: convert.Country(record)}, nil
}

func (s *Service) GetEnterprise(
	ctx context.Context,
	req *geoip_v1_pb.GetEnterpriseRequest,
) (*geoip_v1_pb.GetEnterpriseResponse, error) {
	if err := s.enforceRateLimit(ctx, "GetEnterprise"); err != nil {
		return nil, err
	}

	path, err := s.datasetPath(datasetEnterprise)
	if err != nil {
		return nil, err
	}

	ip, err := parseIP(req.GetIpAddress())
	if err != nil {
		return nil, err
	}

	record, err := withReader(path, func(r *geoip2.Reader) (*geoip2.Enterprise, error) {
		return r.Enterprise(ip)
	})
	if err != nil {
		return nil, err
	}

	return &geoip_v1_pb.GetEnterpriseResponse{Enterprise: convert.Enterprise(record)}, nil
}

func (s *Service) GetAnonymousIp(
	ctx context.Context,
	req *geoip_v1_pb.GetAnonymousIpRequest,
) (*geoip_v1_pb.GetAnonymousIpResponse, error) {
	if err := s.enforceRateLimit(ctx, "GetAnonymousIp"); err != nil {
		return nil, err
	}

	path, err := s.datasetPath(datasetAnonymousIP)
	if err != nil {
		return nil, err
	}

	ip, err := parseIP(req.GetIpAddress())
	if err != nil {
		return nil, err
	}

	record, err := withReader(path, func(r *geoip2.Reader) (*geoip2.AnonymousIP, error) {
		return r.AnonymousIP(ip)
	})
	if err != nil {
		return nil, err
	}

	return &geoip_v1_pb.GetAnonymousIpResponse{AnonymousIp: convert.AnonymousIP(record)}, nil
}

func (s *Service) GetAsn(
	ctx context.Context,
	req *geoip_v1_pb.GetAsnRequest,
) (*geoip_v1_pb.GetAsnResponse, error) {
	if err := s.enforceRateLimit(ctx, "GetAsn"); err != nil {
		return nil, err
	}

	path, err := s.datasetPath(datasetASN)
	if err != nil {
		return nil, err
	}

	ip, err := parseIP(req.GetIpAddress())
	if err != nil {
		return nil, err
	}

	record, err := withReader(path, func(r *geoip2.Reader) (*geoip2.ASN, error) {
		return r.ASN(ip)
	})
	if err != nil {
		return nil, err
	}

	return &geoip_v1_pb.GetAsnResponse{Asn: convert.ASN(record)}, nil
}

func (s *Service) GetConnectionType(
	ctx context.Context,
	req *geoip_v1_pb.GetConnectionTypeRequest,
) (*geoip_v1_pb.GetConnectionTypeResponse, error) {
	if err := s.enforceRateLimit(ctx, "GetConnectionType"); err != nil {
		return nil, err
	}

	path, err := s.datasetPath(datasetConnectionType)
	if err != nil {
		return nil, err
	}

	ip, err := parseIP(req.GetIpAddress())
	if err != nil {
		return nil, err
	}

	record, err := withReader(path, func(r *geoip2.Reader) (*geoip2.ConnectionType, error) {
		return r.ConnectionType(ip)
	})
	if err != nil {
		return nil, err
	}

	return &geoip_v1_pb.GetConnectionTypeResponse{ConnectionType: convert.ConnectionType(record)}, nil
}

func (s *Service) GetDomain(
	ctx context.Context,
	req *geoip_v1_pb.GetDomainRequest,
) (*geoip_v1_pb.GetDomainResponse, error) {
	if err := s.enforceRateLimit(ctx, "GetDomain"); err != nil {
		return nil, err
	}

	path, err := s.datasetPath(datasetDomain)
	if err != nil {
		return nil, err
	}

	ip, err := parseIP(req.GetIpAddress())
	if err != nil {
		return nil, err
	}

	record, err := withReader(path, func(r *geoip2.Reader) (*geoip2.Domain, error) {
		return r.Domain(ip)
	})
	if err != nil {
		return nil, err
	}

	return &geoip_v1_pb.GetDomainResponse{Domain: convert.Domain(record)}, nil
}

func (s *Service) GetIsp(
	ctx context.Context,
	req *geoip_v1_pb.GetIspRequest,
) (*geoip_v1_pb.GetIspResponse, error) {
	if err := s.enforceRateLimit(ctx, "GetIsp"); err != nil {
		return nil, err
	}

	path, err := s.datasetPath(datasetISP)
	if err != nil {
		return nil, err
	}

	ip, err := parseIP(req.GetIpAddress())
	if err != nil {
		return nil, err
	}

	record, err := withReader(path, func(r *geoip2.Reader) (*geoip2.ISP, error) {
		return r.ISP(ip)
	})
	if err != nil {
		return nil, err
	}

	return &geoip_v1_pb.GetIspResponse{Isp: convert.ISP(record)}, nil
}

func (s *Service) enforceRateLimit(ctx context.Context, method string) error {
	if s.limiter == nil || s.rateLimit.IsZero() {
		return nil
	}

	key := method
	if ip, _ := extractClientIP(ctx); ip != "" {
		key = ip + ":" + method
	}

	res, err := s.limiter.Allow(ctx, key, s.rateLimit)
	if err != nil {
		slog.Error("rate limiter failure", "method", method, "error", err)
		return status.Error(codes.Unavailable, "rate limiter unavailable")
	}
	if res.Allowed == 0 {
		return status.Error(codes.ResourceExhausted, "rate limit exceeded")
	}

	return nil
}

func (s *Service) datasetPath(kind string) (string, error) {
	cfg, ok := s.datasets[kind]
	if !ok {
		return "", status.Errorf(codes.Internal, "dataset %s not registered", kind)
	}

	if cfg.path == "" {
		return "", status.Errorf(codes.Unimplemented, "%s dataset not configured", kind)
	}

	if !cfg.available {
		return "", status.Errorf(codes.Unimplemented, "%s dataset unavailable at %q", kind, cfg.path)
	}

	return cfg.path, nil
}

func parseIP(raw string) (netip.Addr, error) {
	ip, err := netip.ParseAddr(raw)
	if err != nil {
		return netip.Addr{}, status.Errorf(codes.InvalidArgument, "invalid ip address %q: %v", raw, err)
	}
	return ip, nil
}

func withReader[T any](path string, fn func(*geoip2.Reader) (T, error)) (T, error) {
	reader, err := geoip2.Open(path)
	if err != nil {
		var zero T
		return zero, err
	}
	defer reader.Close()

	return fn(reader)
}

func fileExists(path string) bool {
	if path == "" {
		return false
	}

	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}
