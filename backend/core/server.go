package core

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"

	"connectrpc.com/connect"
	"github.com/go-chi/chi/v5"
	grpcgwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/encoding/protojson"

	"app/config"
	appv1 "app/domain/service/proto/app/v1"
	"app/gen/buf/app/v1/appv1connect"
)

var (
	jsonCodec = &JSONCodec{
		grpcgwruntime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: true,
			},
		},
	}

	connectCORSOptions = cors.Options{
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
	}
)

type JSONCodec struct {
	grpcgwruntime.JSONPb
}

func (codec JSONCodec) Name() string {
	return "json"
}

type Server struct {
	http.Server

	config config.ServerConfig
}

func NewServer(conf config.ServerConfig, taskService *appv1.TaskService) *Server {
	var grpcHandler http.Handler
	{
		r := chi.NewRouter()
		{
			path, h := appv1connect.NewTaskServiceHandler(
				taskService,
				connect.WithCodec(jsonCodec),
				connect.WithInterceptors(
					newErrorHandlingInterceptor(),
					newValidationInterceptor(),
				),
			)

			r.Handle(path+"*", h)
		}

		grpcHandler = r
		grpcHandler = cors.New(connectCORSOptions).Handler(grpcHandler)
		grpcHandler = h2c.NewHandler(grpcHandler, &http2.Server{})
	}

	r := chi.NewRouter()
	r.Mount("/grpc", http.StripPrefix("/grpc", grpcHandler))

	return &Server{
		Server: http.Server{
			Addr:    conf.Addr(),
			Handler: r,
		},

		config: conf,
	}
}

func newErrorHandlingInterceptor() connect.Interceptor {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (resp connect.AnyResponse, err error) {
			unknownErr := connect.NewError(connect.CodeUnknown, errors.New("unknown error occured"))

			defer func() {
				if v := recover(); v != nil {
					e, ok := v.(error)
					if !ok {
						e = fmt.Errorf("%v", v)
					}

					slog.LogAttrs(ctx, slog.LevelError, e.Error(), slog.String("stack", string(debug.Stack())))
					err = unknownErr
				}
			}()

			resp, err = next(ctx, req)
			if err != nil {
				var connectErr *connect.Error
				if errors.As(err, &connectErr) {
					if connectErr.Code() == connect.CodeUnknown {
						slog.Error(connectErr.Message())
						err = unknownErr
					}
				} else {
					slog.Error(err.Error())
					err = unknownErr
				}
			}

			return
		})
	})
}

func newValidationInterceptor() connect.Interceptor {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
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
	})
}
