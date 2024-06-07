package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

const (
	// grpc info
	grpcGatewayUserAgentHeader = "user-agent"
	grpcXForwaredHostHeader    = ":authority"

	// gateway info
	gatewayUserAgentHeader     = "grpcgateway-user-agent"
	gatewayXForwatedHostHeader = "x-forwarded-host"
)

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mdtd := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			mdtd.UserAgent = userAgents[0]
		}

		if clientIPs := md.Get(grpcXForwaredHostHeader); len(clientIPs) > 0 {
			mdtd.ClientIP = clientIPs[0]
		}
		if userAgents := md.Get(gatewayUserAgentHeader); len(userAgents) > 0 {
			mdtd.UserAgent = userAgents[0]
		}

		if clientIPs := md.Get(gatewayXForwatedHostHeader); len(clientIPs) > 0 {
			mdtd.ClientIP = clientIPs[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		mdtd.ClientIP = p.Addr.String()
	}

	return mdtd
}
