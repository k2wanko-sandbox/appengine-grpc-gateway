package internal

import (
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	ctxKeyConn       = "grpc.ClientConn key"
	ctxKeyEchoClient = "EchoClient key"
)

type newClientFunc func(*grpc.ClientConn) (interface{}, interface{})

var (
	newClientFuncsMu sync.Mutex
	newClientFuncs   = make(map[string]newClientFunc)
)

func RegisterNewClientFunc(name string, f newClientFunc) {
	newClientFuncsMu.Lock()
	defer newClientFuncsMu.Unlock()
	if _, dup := newClientFuncs[name]; dup {
		panic("gateway/internal: Register called twice for NewClientFunc " + name)
	}
	newClientFuncs[name] = f
}

func registerClients(ctx context.Context, conn *grpc.ClientConn) context.Context {
	for _, f := range newClientFuncs {
		k, c := f(conn)
		ctx = context.WithValue(ctx, k, c)
	}
	return ctx
}

func NewContext(conn *grpc.ClientConn) context.Context {
	return WithContext(context.Background(), conn)
}

func WithContext(ctx context.Context, conn *grpc.ClientConn) context.Context {
	if conn := Conn(ctx); conn != nil {
		return ctx
	}
	ctx = registerClients(ctx, conn)
	return context.WithValue(ctx, &ctxKeyConn, conn)
}

func Conn(ctx context.Context) *grpc.ClientConn {
	if conn, ok := ctx.Value(&ctxKeyConn).(*grpc.ClientConn); ok {
		return conn
	}
	return nil
}
