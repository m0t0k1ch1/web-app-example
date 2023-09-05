package main

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/encoding/protojson"

	"app/config"
	"app/gen/buf/app/v1/appv1connect"
	appv1 "app/service/app/v1"
)

type jsonCodec struct {
	runtime.JSONPb
}

func (codec jsonCodec) Name() string {
	return "json"
}

type Server struct {
	http.Server
}

func NewServer(conf config.Server, taskService *appv1.TaskService) *Server {
	var grpcHandler http.Handler
	{
		r := chi.NewRouter()

		path, h := appv1connect.NewTaskServiceHandler(
			taskService,
			connect.WithInterceptors(
				connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
					return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
						v, ok := req.Any().(interface {
							ValidateAll() error
						})
						if ok {
							if err := v.ValidateAll(); err != nil {
								return nil, connect.NewError(connect.CodeInvalidArgument, err)
							}
						}

						return next(ctx, req)
					})
				}),
			),
			connect.WithCodec(&jsonCodec{
				runtime.JSONPb{
					MarshalOptions: protojson.MarshalOptions{
						EmitUnpopulated: true,
					},
				},
			}),
		)

		c := cors.New(cors.Options{
			AllowedMethods: []string{
				http.MethodGet,
				http.MethodPost,
			},
			AllowedHeaders: []string{
				"Accept-Encoding",
				"Content-Encoding",
				"Content-Type",
				"Connect-Protocol-Version",
				"Connect-Timeout-Ms",
				"Connect-Accept-Encoding",  // Unused in web browsers, but added for future-proofing
				"Connect-Content-Encoding", // Unused in web browsers, but added for future-proofing
				"Grpc-Timeout",             // Used for gRPC-web
				"X-Grpc-Web",               // Used for gRPC-web
				"X-User-Agent",             // Used for gRPC-web
			},
			ExposedHeaders: []string{
				"Content-Encoding",         // Unused in web browsers, but added for future-proofing
				"Connect-Content-Encoding", // Unused in web browsers, but added for future-proofing
				"Grpc-Status",              // Required for gRPC-web
				"Grpc-Message",             // Required for gRPC-web
			},
		})

		r.Handle(path+"*", h2c.NewHandler(c.Handler(h), &http2.Server{}))

		grpcHandler = r
	}

	r := chi.NewRouter()
	r.Mount("/grpc", http.StripPrefix("/grpc", grpcHandler))

	return &Server{
		Server: http.Server{
			Addr:    conf.Addr(),
			Handler: r,
		},
	}
}
