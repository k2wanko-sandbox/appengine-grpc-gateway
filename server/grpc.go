package main

import (
	"net/http"

	"github.com/k2wanko-sandbox/appengine-grpc-gateway/echo"
	"github.com/k2wanko/grpc-pipe/gateway"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/grpc"
)

var (
	s *gateway.Server
)

func appCtxInjector(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	r := s.Request(ctx)
	ctx = appengine.WithContext(ctx, r)
	return handler(ctx, req)
}

func createGRPCHandler(ctx context.Context) http.Handler {
	s = gateway.New(ctx,
		gateway.WithGrpcOptions(grpc.UnaryInterceptor(appCtxInjector)))
	echo.RegisterService(s)
	return s
}
