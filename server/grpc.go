package main

import (
	"net/http"
	"sync"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/k2wanko-sandbox/appengine-grpc-gateway/echo"
	"github.com/k2wanko/grpc-pipe"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	reqID = "x-appengine-request-id-hash"
	reqMu sync.RWMutex
	reqs  = make(map[string]*http.Request)
)

func reqFromCtx(ctx context.Context) *http.Request {
	md, _ := metadata.FromContext(ctx)
	if md == nil {
		return nil
	}
	ids := md[reqID]
	if len(ids) < 1 {
		return nil
	}
	id := ids[0]
	reqMu.RLock()
	defer reqMu.RUnlock()
	if req, ok := reqs[id]; ok {
		return req
	}
	return nil
}

func appCtxInjector(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	r := reqFromCtx(ctx)
	ctx = appengine.WithContext(ctx, r)
	return handler(ctx, req)
}

func createGRPCHandler(ctx context.Context) http.Handler {
	l := pipe.Listen()
	s := grpc.NewServer(grpc.UnaryInterceptor(appCtxInjector))
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for name, hs := range r.Header {
			for _, h := range hs {
				r.Header.Add("Grpc-Metadata-"+name, h)
			}
		}
		key := r.Header.Get("X-Appengine-Request-Id-Hash")
		reqMu.Lock()
		reqs[key] = r
		reqMu.Unlock()
		defer func() {
			reqMu.Lock()
			defer reqMu.Unlock()
			delete(reqs, key)
		}()
		mux.ServeHTTP(w, r)
	})
}
