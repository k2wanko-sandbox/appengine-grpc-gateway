// +build !appengine

package main

import (
	"net"
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/k2wanko-sandbox/appengine-grpc-gateway/echo"
)

func main() {
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	echo.RegisterServer(s)
	go s.Serve(l)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		cancel()
		s.Stop()
	}()

	mux := runtime.NewServeMux()
	err = echo.RegisterGateway(ctx, mux, "localhost:9090", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	http.Handle("/", mux)
	http.ListenAndServe(":8080", nil)
}
