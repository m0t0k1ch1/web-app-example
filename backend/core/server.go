package core

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"
	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	grpcgwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/encoding/protojson"

	"app/config"
	"app/gen/buf/app/v1/appv1connect"
	appv1 "app/service/app/v1"
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

func NewServer(conf config.ServerConfig, sentryHandler *sentryhttp.Handler, taskService *appv1.TaskService) *Server {
	var grpcHandler http.Handler
	{
		r := chi.NewRouter()
		{
			path, h := appv1connect.NewTaskServiceHandler(
				taskService,
				connect.WithCodec(jsonCodec),
				connect.WithInterceptors(
					errorReportInterceptor(),
					validationInterceptor(),
				),
			)

			r.Handle(path+"*", h)
		}

		grpcHandler = r
		grpcHandler = cors.New(connectCORSOptions).Handler(grpcHandler)
		grpcHandler = h2c.NewHandler(grpcHandler, &http2.Server{})
	}

	var testHandler http.Handler
	{
		r := chi.NewRouter()
		{
			r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
				panic("y tho")
			})
		}

		testHandler = r
	}

	r := chi.NewRouter()
	r.Use(chimiddleware.Recoverer)
	if sentryHandler != nil {
		r.Use(func(h http.Handler) http.Handler {
			return sentryHandler.Handle(h)
		})
	}
	r.Mount("/grpc", http.StripPrefix("/grpc", grpcHandler))
	r.Mount("/test", http.StripPrefix("/test", testHandler))

	return &Server{
		Server: http.Server{
			Addr:    conf.Addr(),
			Handler: r,
		},

		config: conf,
	}
}

func errorReportInterceptor() connect.Interceptor {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			resp, err := next(ctx, req)
			if err != nil {
				if hub := sentry.GetHubFromContext(ctx); hub != nil {
					event, _ := errors.BuildSentryReport(err)
					hub.CaptureEvent(event)
				}
			}

			return resp, err
		})
	})
}

func validationInterceptor() connect.Interceptor {
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
