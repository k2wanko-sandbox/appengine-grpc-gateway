package echo

import (
	"google.golang.org/appengine/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"golang.org/x/net/context"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/k2wanko-sandbox/appengine-grpc-gateway/internal"
	pb "github.com/k2wanko-sandbox/appengine-grpc-gateway/internal/echo"
	"github.com/k2wanko/grpc-pipe/gateway"
)

var key = "echo service client key"

func newClient(conn *grpc.ClientConn) (interface{}, interface{}) {
	return &key, pb.NewEchoServiceClient(conn)
}

func client(ctx context.Context) pb.EchoServiceClient {
	if c, ok := ctx.Value(&key).(pb.EchoServiceClient); ok {
		return c
	}
	return nil
}

func init() {
	internal.RegisterNewClientFunc("EchoService", newClient)
}

func Echo(ctx context.Context, msg string) (string, error) {
	m := &pb.Message{Value: msg}
	res, err := client(ctx).Echo(ctx, m)
	if err != nil {
		return "", err
	}
	return res.Value, nil
}

// Gateway

func RegisterGateway(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return pb.RegisterEchoServiceHandler(ctx, mux, conn)
}

// Server

func RegisterServer(srv *grpc.Server) {
	pb.RegisterEchoServiceServer(srv, new(Service))
}

var _ pb.EchoServiceServer = &Service{}

type Service struct{}

func (s *Service) Echo(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	md, _ := metadata.FromContext(ctx)
	log.Infof(ctx, "Metadata: %#v", md)
	msg.Value = "Server: " + msg.Value
	return msg, nil
}

// grpc-pipe
func RegisterService(s *gateway.Server) {
	s.RegisterService(pb.RegisterEchoServiceServer, pb.RegisterEchoServiceHandler, new(Service))
}
