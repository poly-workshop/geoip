package configs

import (
	"github.com/poly-workshop/go-webmods/app"
)

type GRPCServerConfig struct {
	Port                 uint
	CityDBPath           string
	CountryDBPath        string
	EnterpriseDBPath     string
	AnonymousIPDBPath    string
	ASNDBPath            string
	ConnectionTypeDBPath string
	DomainDBPath         string
	ISPDBPath            string
	RedisURLs            []string
	RedisPassword        string
	RateLimitPerMinute   int
	RateLimitBurst       int
	MaxMindLicenseKey    string
	MaxMindAccountID     string
}

type GatewayConfig struct {
	Port         uint
	GRPCEndpoint string
}

func LoadGRPCServerConfig() GRPCServerConfig {
	cfg := app.Config()

	return GRPCServerConfig{
		Port:                 cfg.GetUint(PortConfigKey),
		CityDBPath:           cfg.GetString(CityDBPathConfigKey),
		CountryDBPath:        cfg.GetString(CountryDBPathConfigKey),
		EnterpriseDBPath:     cfg.GetString(EnterpriseDBPathConfigKey),
		AnonymousIPDBPath:    cfg.GetString(AnonymousIPDBPathConfigKey),
		ASNDBPath:            cfg.GetString(ASNDBPathConfigKey),
		ConnectionTypeDBPath: cfg.GetString(ConnectionTypeDBPathConfigKey),
		DomainDBPath:         cfg.GetString(DomainDBPathConfigKey),
		ISPDBPath:            cfg.GetString(ISPDBPathConfigKey),
		RedisURLs:            cfg.GetStringSlice(RedisURLsConfigKey),
		RedisPassword:        cfg.GetString(RedisPasswordConfigKey),
		RateLimitPerMinute:   cfg.GetInt(RateLimitPerMinuteConfigKey),
		RateLimitBurst:       cfg.GetInt(RateLimitBurstConfigKey),
		MaxMindLicenseKey:    cfg.GetString(MaxMindLicenseKeyConfigKey),
		MaxMindAccountID:     cfg.GetString(MaxMindAccountIDConfigKey),
	}
}

func LoadGatewayConfig() GatewayConfig {
	cfg := app.Config()

	return GatewayConfig{
		Port:         cfg.GetUint(PortConfigKey),
		GRPCEndpoint: cfg.GetString(GRPCEndpointKey),
	}
}
