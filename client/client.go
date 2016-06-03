package main

import (
	"fmt"
	"log"

	"github.com/k2wanko-sandbox/appengine-grpc-gateway"
	"github.com/k2wanko-sandbox/appengine-grpc-gateway/echo"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ctx := gateway.NewContext(conn)

	msg := "Hello"
	fmt.Printf("Send: %s\n", msg)
	res, err := echo.Echo(ctx, msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server: %s\n", res)

}
