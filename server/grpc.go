package main

import (
	"net/http"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/k2wanko-sandbox/appengine-grpc-gateway/echo"
	"github.com/k2wanko/grpc-pipe"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func createGRPCHandler(ctx context.Context) http.Handler {
	l := pipe.Listen()
	s := grpc.NewServer()
	echo.RegisterServer(s)
	go s.Serve(l)

	conn, err := grpc.Dial("", grpc.WithInsecure(), l.WithDialer())
	if err != nil {
		panic(err)
	}

	mux := runtime.NewServeMux()
	err = echo.RegisterGateway(ctx, mux, conn)
	if err != nil {
		panic(err)
	}
	return mux
}
