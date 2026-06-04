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
	grpcGatewayUserAgent = "grpcgateway-user-agent"
	userAgentHeader      = "user-agent"
	xForwarderForHeader  = "x-forwarded-for"
)

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return mtdt
	}

	if userAgent := md.Get(grpcGatewayUserAgent); len(userAgent) > 0 {
		mtdt.UserAgent = userAgent[0]
	}

	if userAgent := md.Get(userAgentHeader); len(userAgent) > 0 {
		mtdt.UserAgent = userAgent[0]
	}

	if clientIP := md.Get(xForwarderForHeader); len(clientIP) > 0 {
		mtdt.ClientIP = clientIP[0]
	}

	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIP = p.Addr.String()
	}

	return mtdt
}
