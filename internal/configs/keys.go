package configs

const (
	PortConfigKey                 = "port"
	GRPCEndpointKey               = "grpc_endpoint"
	CityDBPathConfigKey           = "geoip2.databases.city"
	CountryDBPathConfigKey        = "geoip2.databases.country"
	EnterpriseDBPathConfigKey     = "geoip2.databases.enterprise"
	AnonymousIPDBPathConfigKey    = "geoip2.databases.anonymous_ip"
	ASNDBPathConfigKey            = "geoip2.databases.asn"
	ConnectionTypeDBPathConfigKey = "geoip2.databases.connection_type"
	DomainDBPathConfigKey         = "geoip2.databases.domain"
	ISPDBPathConfigKey            = "geoip2.databases.isp"
	RedisURLsConfigKey            = "redis.urls"
	RedisPasswordConfigKey        = "redis.password"
	RateLimitPerMinuteConfigKey   = "rate_limit.per_minute"
	RateLimitBurstConfigKey       = "rate_limit.burst"
	MaxMindLicenseKeyConfigKey    = "maxmind.license_key"
	MaxMindAccountIDConfigKey     = "maxmind.account_id"
)
