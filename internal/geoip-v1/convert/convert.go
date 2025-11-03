package convert

import (
	geoip2 "github.com/oschwald/geoip2-golang/v2"
	geoip_v1_pb "github.com/poly-workshop/geoip/gen/go/geoip/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Names converts geoip2 localized names to their protobuf representation.
func Names(src geoip2.Names) *geoip_v1_pb.Names {
	if !src.HasData() {
		return nil
	}

	return &geoip_v1_pb.Names{
		De:   src.German,
		En:   src.English,
		Es:   src.Spanish,
		Fr:   src.French,
		Ja:   src.Japanese,
		PtBr: src.BrazilianPortuguese,
		Ru:   src.Russian,
		ZhCn: src.SimplifiedChinese,
	}
}

// City converts a City database record to protobuf.
func City(src *geoip2.City) *geoip_v1_pb.City {
	if src == nil || !src.HasData() {
		return nil
	}

	return &geoip_v1_pb.City{
		CityRecord:         CityRecord(src.City),
		Continent:          Continent(src.Continent),
		Country:            CountryRecord(src.Country),
		RegisteredCountry:  CountryRecord(src.RegisteredCountry),
		RepresentedCountry: RepresentedCountry(src.RepresentedCountry),
		Location:           Location(src.Location),
		Postal:             CityPostal(src.Postal),
		Subdivisions:       CitySubdivisions(src.Subdivisions),
		Traits:             CityTraits(src.Traits),
	}
}

// Country converts a Country database record to protobuf.
func Country(src *geoip2.Country) *geoip_v1_pb.Country {
	if src == nil || !src.HasData() {
		return nil
	}

	return &geoip_v1_pb.Country{
		Continent:          Continent(src.Continent),
		Country:            CountryRecord(src.Country),
		RegisteredCountry:  CountryRecord(src.RegisteredCountry),
		RepresentedCountry: RepresentedCountry(src.RepresentedCountry),
		Traits:             CountryTraits(src.Traits),
	}
}

// Enterprise converts an Enterprise database record to protobuf.
func Enterprise(src *geoip2.Enterprise) *geoip_v1_pb.Enterprise {
	if src == nil || !src.HasData() {
		return nil
	}

	return &geoip_v1_pb.Enterprise{
		Continent:          Continent(src.Continent),
		Subdivisions:       EnterpriseSubdivisions(src.Subdivisions),
		Postal:             EnterprisePostal(src.Postal),
		RepresentedCountry: RepresentedCountry(src.RepresentedCountry),
		Country:            EnterpriseCountryRecord(src.Country),
		RegisteredCountry:  CountryRecord(src.RegisteredCountry),
		City:               EnterpriseCityRecord(src.City),
		Location:           Location(src.Location),
		Traits:             EnterpriseTraits(src.Traits),
	}
}

// AnonymousIP converts AnonymousIP database record.
func AnonymousIP(src *geoip2.AnonymousIP) *geoip_v1_pb.AnonymousIp {
	if src == nil || !src.HasData() {
		return nil
	}

	ip := ""
	if src.IPAddress.IsValid() {
		ip = src.IPAddress.String()
	}

	network := ""
	if src.Network.IsValid() {
		network = src.Network.String()
	}

	return &geoip_v1_pb.AnonymousIp{
		IpAddress:          ip,
		Network:            network,
		IsAnonymous:        src.IsAnonymous,
		IsAnonymousVpn:     src.IsAnonymousVPN,
		IsHostingProvider:  src.IsHostingProvider,
		IsPublicProxy:      src.IsPublicProxy,
		IsResidentialProxy: src.IsResidentialProxy,
		IsTorExitNode:      src.IsTorExitNode,
	}
}

// ASN converts ASN database record.
func ASN(src *geoip2.ASN) *geoip_v1_pb.Asn {
	if src == nil || !src.HasData() {
		return nil
	}

	ip := ""
	if src.IPAddress.IsValid() {
		ip = src.IPAddress.String()
	}

	network := ""
	if src.Network.IsValid() {
		network = src.Network.String()
	}

	return &geoip_v1_pb.Asn{
		IpAddress:                    ip,
		Network:                      network,
		AutonomousSystemNumber:       uint64(src.AutonomousSystemNumber),
		AutonomousSystemOrganization: src.AutonomousSystemOrganization,
	}
}

// ConnectionType converts ConnectionType record.
func ConnectionType(src *geoip2.ConnectionType) *geoip_v1_pb.ConnectionType {
	if src == nil || !src.HasData() {
		return nil
	}

	ip := ""
	if src.IPAddress.IsValid() {
		ip = src.IPAddress.String()
	}

	network := ""
	if src.Network.IsValid() {
		network = src.Network.String()
	}

	return &geoip_v1_pb.ConnectionType{
		ConnectionType: src.ConnectionType,
		IpAddress:      ip,
		Network:        network,
	}
}

// Domain converts Domain record.
func Domain(src *geoip2.Domain) *geoip_v1_pb.Domain {
	if src == nil || !src.HasData() {
		return nil
	}

	ip := ""
	if src.IPAddress.IsValid() {
		ip = src.IPAddress.String()
	}

	network := ""
	if src.Network.IsValid() {
		network = src.Network.String()
	}

	return &geoip_v1_pb.Domain{
		Domain:    src.Domain,
		IpAddress: ip,
		Network:   network,
	}
}

// ISP converts ISP record.
func ISP(src *geoip2.ISP) *geoip_v1_pb.Isp {
	if src == nil || !src.HasData() {
		return nil
	}

	ip := ""
	if src.IPAddress.IsValid() {
		ip = src.IPAddress.String()
	}

	network := ""
	if src.Network.IsValid() {
		network = src.Network.String()
	}

	return &geoip_v1_pb.Isp{
		Network:                      network,
		IpAddress:                    ip,
		AutonomousSystemOrganization: src.AutonomousSystemOrganization,
		Isp:                          src.ISP,
		MobileCountryCode:            src.MobileCountryCode,
		MobileNetworkCode:            src.MobileNetworkCode,
		Organization:                 src.Organization,
		AutonomousSystemNumber:       uint64(src.AutonomousSystemNumber),
	}
}

// CityRecord converts city record details.
func CityRecord(src geoip2.CityRecord) *geoip_v1_pb.CityRecord {
	if !src.HasData() {
		return nil
	}

	return &geoip_v1_pb.CityRecord{
		GeonameId: uint64(src.GeoNameID),
		Names:     Names(src.Names),
	}
}

// Continent converts continent data.
func Continent(src geoip2.Continent) *geoip_v1_pb.Continent {
	if !src.HasData() {
		return nil
	}

	return &geoip_v1_pb.Continent{
		Names:     Names(src.Names),
		Code:      src.Code,
		GeonameId: uint64(src.GeoNameID),
	}
}

// CountryRecord converts country record details.
func CountryRecord(src geoip2.CountryRecord) *geoip_v1_pb.CountryRecord {
	if !src.HasData() {
		return nil
	}

	return &geoip_v1_pb.CountryRecord{
		Names:             Names(src.Names),
		IsoCode:           src.ISOCode,
		GeonameId:         uint64(src.GeoNameID),
		IsInEuropeanUnion: src.IsInEuropeanUnion,
	}
}

// RepresentedCountry converts represented country details.
func RepresentedCountry(src geoip2.RepresentedCountry) *geoip_v1_pb.RepresentedCountry {
	if !src.HasData() {
		return nil
	}

	return &geoip_v1_pb.RepresentedCountry{
		Names:             Names(src.Names),
		IsoCode:           src.ISOCode,
		GeonameId:         uint64(src.GeoNameID),
		IsInEuropeanUnion: src.IsInEuropeanUnion,
		Type:              src.Type,
	}
}

// Location converts location details.
func Location(src geoip2.Location) *geoip_v1_pb.Location {
	if !src.HasData() {
		return nil
	}

	out := &geoip_v1_pb.Location{
		TimeZone:       src.TimeZone,
		MetroCode:      uint32(src.MetroCode),
		AccuracyRadius: uint32(src.AccuracyRadius),
	}

	if src.Latitude != nil {
		out.Latitude = wrapperspb.Double(*src.Latitude)
	}
	if src.Longitude != nil {
		out.Longitude = wrapperspb.Double(*src.Longitude)
	}

	return out
}

// CityPostal converts postal code data.
func CityPostal(src geoip2.CityPostal) *geoip_v1_pb.CityPostal {
	if !src.HasData() {
		return nil
	}

	return &geoip_v1_pb.CityPostal{Code: src.Code}
}

// CitySubdivisions converts subdivision data for city records.
func CitySubdivisions(src []geoip2.CitySubdivision) []*geoip_v1_pb.CitySubdivision {
	if len(src) == 0 {
		return nil
	}

	var out []*geoip_v1_pb.CitySubdivision
	for _, subdivision := range src {
		if !subdivision.HasData() {
			continue
		}

		out = append(out, &geoip_v1_pb.CitySubdivision{
			Names:     Names(subdivision.Names),
			IsoCode:   subdivision.ISOCode,
			GeonameId: uint64(subdivision.GeoNameID),
		})
	}

	if len(out) == 0 {
		return nil
	}

	return out
}

// CityTraits converts trait data for city records.
func CityTraits(src geoip2.CityTraits) *geoip_v1_pb.CityTraits {
	if !src.HasData() {
		return nil
	}

	ip := ""
	if src.IPAddress.IsValid() {
		ip = src.IPAddress.String()
	}

	network := ""
	if src.Network.IsValid() {
		network = src.Network.String()
	}

	return &geoip_v1_pb.CityTraits{
		IsAnycast: src.IsAnycast,
		IpAddress: ip,
		Network:   network,
	}
}

// CountryTraits converts trait data for country records.
func CountryTraits(src geoip2.CountryTraits) *geoip_v1_pb.CountryTraits {
	if !src.HasData() {
		return nil
	}

	ip := ""
	if src.IPAddress.IsValid() {
		ip = src.IPAddress.String()
	}

	network := ""
	if src.Network.IsValid() {
		network = src.Network.String()
	}

	return &geoip_v1_pb.CountryTraits{
		IsAnycast: src.IsAnycast,
		IpAddress: ip,
		Network:   network,
	}
}

// EnterpriseCityRecord converts enterprise city data.
func EnterpriseCityRecord(src geoip2.EnterpriseCityRecord) *geoip_v1_pb.EnterpriseCityRecord {
	if !src.HasData() {
		return nil
	}

	return &geoip_v1_pb.EnterpriseCityRecord{
		Names:      Names(src.Names),
		GeonameId:  uint64(src.GeoNameID),
		Confidence: uint32(src.Confidence),
	}
}

// EnterprisePostal converts enterprise postal data.
func EnterprisePostal(src geoip2.EnterprisePostal) *geoip_v1_pb.EnterprisePostal {
	if !src.HasData() {
		return nil
	}

	return &geoip_v1_pb.EnterprisePostal{
		Code:       src.Code,
		Confidence: uint32(src.Confidence),
	}
}

// EnterpriseSubdivisions converts enterprise subdivision data.
func EnterpriseSubdivisions(src []geoip2.EnterpriseSubdivision) []*geoip_v1_pb.EnterpriseSubdivision {
	if len(src) == 0 {
		return nil
	}

	var out []*geoip_v1_pb.EnterpriseSubdivision
	for _, subdivision := range src {
		if !subdivision.HasData() {
			continue
		}

		out = append(out, &geoip_v1_pb.EnterpriseSubdivision{
			Names:      Names(subdivision.Names),
			IsoCode:    subdivision.ISOCode,
			GeonameId:  uint64(subdivision.GeoNameID),
			Confidence: uint32(subdivision.Confidence),
		})
	}

	if len(out) == 0 {
		return nil
	}

	return out
}

// EnterpriseCountryRecord converts enterprise country data.
func EnterpriseCountryRecord(src geoip2.EnterpriseCountryRecord) *geoip_v1_pb.EnterpriseCountryRecord {
	if !src.HasData() {
		return nil
	}

	return &geoip_v1_pb.EnterpriseCountryRecord{
		Names:             Names(src.Names),
		IsoCode:           src.ISOCode,
		GeonameId:         uint64(src.GeoNameID),
		Confidence:        uint32(src.Confidence),
		IsInEuropeanUnion: src.IsInEuropeanUnion,
	}
}

// EnterpriseTraits converts enterprise traits data.
func EnterpriseTraits(src geoip2.EnterpriseTraits) *geoip_v1_pb.EnterpriseTraits {
	if !src.HasData() {
		return nil
	}

	ip := ""
	if src.IPAddress.IsValid() {
		ip = src.IPAddress.String()
	}

	network := ""
	if src.Network.IsValid() {
		network = src.Network.String()
	}

	return &geoip_v1_pb.EnterpriseTraits{
		Network:                      network,
		IpAddress:                    ip,
		AutonomousSystemOrganization: src.AutonomousSystemOrganization,
		ConnectionType:               src.ConnectionType,
		Domain:                       src.Domain,
		Isp:                          src.ISP,
		MobileCountryCode:            src.MobileCountryCode,
		MobileNetworkCode:            src.MobileNetworkCode,
		Organization:                 src.Organization,
		UserType:                     src.UserType,
		StaticIpScore:                src.StaticIPScore,
		AutonomousSystemNumber:       uint64(src.AutonomousSystemNumber),
		IsAnycast:                    src.IsAnycast,
		IsLegitimateProxy:            src.IsLegitimateProxy,
	}
}
