// +build !appengine

package main

import (
	"net"

	"google.golang.org/grpc"

	"github.com/k2wanko-sandbox/appengine-grpc-gateway/echo"
	pb "github.com/k2wanko-sandbox/appengine-grpc-gateway/internal/echo"
)

func main() {
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, new(echo.Service))
	s.Serve(l)
}
