package core

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	graphqlerrcode "github.com/99designs/gqlgen/graphql/errcode"
	graphqlhandler "github.com/99designs/gqlgen/graphql/handler"
	graphqlplayground "github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/samber/oops"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"app/config"
	"app/domain/e"
	"app/domain/resolver"
	"app/gen/gqlgen"
)

var (
	corsOptions = cors.Options{
		AllowedOrigins: []string{
			"*",
		},
		AllowedMethods: []string{
			http.MethodPost,
		},
		AllowedHeaders: []string{
			"Content-Type",
		},
	}
)

func init() {
	graphqlerrcode.RegisterErrorType(graphqlerrcode.ValidationFailed, graphqlerrcode.KindUser)
	graphqlerrcode.RegisterErrorType(graphqlerrcode.ParseFailed, graphqlerrcode.KindUser)
}

type Server struct {
	http.Server

	config config.ServerConfig
}

func NewServer(
	srvConf config.ServerConfig,
	resolver *resolver.Resolver,
) *Server {
	var gqlHandler http.Handler
	{
		r := chi.NewRouter()
		{
			h := graphqlhandler.NewDefaultServer(gqlgen.NewExecutableSchema(gqlgen.Config{
				Resolvers: resolver,
			}))
			{
				h.SetRecoverFunc(func(ctx context.Context, panicErr any) error {
					err, ok := panicErr.(error)
					if !ok {
						err = fmt.Errorf("%v", panicErr)
					}

					return err
				})

				h.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
					var (
						isUnexpectedErr bool = false
						gqlErr          *gqlerror.Error
					)
					{
						if errors.As(err, &gqlErr) {
							code, ok := gqlErr.Extensions["code"]
							if ok {
								if code == gqlgen.ErrorCodeInternalServerError {
									isUnexpectedErr = true
								}
							} else {
								isUnexpectedErr = true
								gqlErr.Extensions["code"] = gqlgen.ErrorCodeInternalServerError
							}
						} else {
							isUnexpectedErr = true
							gqlErr = e.NewUnexpectedGQLError(ctx, err)
						}
					}

					if isUnexpectedErr {
						slog.ErrorContext(ctx, err.Error(), slog.Any("error", oops.Wrap(err)))
					}

					return gqlErr
				})
			}

			r.Use(cors.New(corsOptions).Handler)
			r.Handle("/", h)
		}

		gqlHandler = r
	}

	r := chi.NewRouter()
	r.Handle("/graphql", http.StripPrefix("/graphql", gqlHandler))

	if srvConf.WithPlayground {
		r.Handle("/", graphqlplayground.Handler("GraphQL playground", "/graphql"))
	}

	return &Server{
		Server: http.Server{
			Addr:    srvConf.Addr(),
			Handler: r,
		},

		config: srvConf,
	}
}
