package geoip_v1

import (
	"context"
	"net"
	"strings"

	geoip_v1_pb "github.com/poly-workshop/geoip/gen/go/geoip/v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

// extractClientIP attempts to determine the real client IP address from gRPC context.
// It checks for common proxy headers first, then falls back to the peer address.
func extractClientIP(ctx context.Context) (ip, source string) {
	// Check for HTTP headers in gRPC metadata (from grpc-gateway)
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		// Check X-Forwarded-For header (most common)
		if xff := md.Get("x-forwarded-for"); len(xff) > 0 {
			// X-Forwarded-For can contain multiple IPs: "client, proxy1, proxy2"
			// The first IP is typically the original client
			ips := strings.Split(xff[0], ",")
			if len(ips) > 0 {
				clientIP := strings.TrimSpace(ips[0])
				if clientIP != "" {
					return clientIP, "x-forwarded-for"
				}
			}
		}

		// Check X-Real-IP header (common with nginx)
		if xri := md.Get("x-real-ip"); len(xri) > 0 {
			if xri[0] != "" {
				return xri[0], "x-real-ip"
			}
		}

		// Check CF-Connecting-IP header (Cloudflare)
		if cfip := md.Get("cf-connecting-ip"); len(cfip) > 0 {
			if cfip[0] != "" {
				return cfip[0], "cf-connecting-ip"
			}
		}

		// Check X-Original-Forwarded-For header
		if xoff := md.Get("x-original-forwarded-for"); len(xoff) > 0 {
			if xoff[0] != "" {
				return xoff[0], "x-original-forwarded-for"
			}
		}

		// Check True-Client-IP header (Akamai)
		if tcip := md.Get("true-client-ip"); len(tcip) > 0 {
			if tcip[0] != "" {
				return tcip[0], "true-client-ip"
			}
		}
	}

	// Fall back to peer address from gRPC context
	if p, ok := peer.FromContext(ctx); ok && p.Addr != nil {
		if tcpAddr, ok := p.Addr.(*net.TCPAddr); ok {
			return tcpAddr.IP.String(), "remote_addr"
		}
		// Handle other address types
		host, _, err := net.SplitHostPort(p.Addr.String())
		if err == nil {
			return host, "remote_addr"
		}
		return p.Addr.String(), "remote_addr"
	}

	return "", "unknown"
}

func (s *Service) GetMyIp(
	ctx context.Context,
	req *geoip_v1_pb.GetMyIpRequest,
) (*geoip_v1_pb.GetMyIpResponse, error) {
	ip, source := extractClientIP(ctx)

	return &geoip_v1_pb.GetMyIpResponse{
		IpAddress:    ip,
		DetectedFrom: source,
	}, nil
}
