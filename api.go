package gateway

import (
	"github.com/k2wanko-sandbox/appengine-grpc-gateway/internal"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func NewContext(conn *grpc.ClientConn) context.Context {
	return WithContext(context.Background(), conn)
}

func WithContext(ctx context.Context, conn *grpc.ClientConn) context.Context {
	return internal.WithContext(ctx, conn)
}
