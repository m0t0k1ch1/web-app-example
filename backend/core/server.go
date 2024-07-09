package core

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	graphqlerrcode "github.com/99designs/gqlgen/graphql/errcode"
	graphqlhandler "github.com/99designs/gqlgen/graphql/handler"
	graphqlplayground "github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/samber/oops"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"app/config"
	"app/gen/gqlgen"
	"app/gql"
	"app/gql/errcode"
)

var (
	corsOptions = cors.Options{
		AllowedOrigins: []string{
			"*",
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
	gqlResolver *gql.Resolver,
) *Server {
	var gqlHandler http.Handler
	{
		r := chi.NewRouter()
		{
			h := graphqlhandler.NewDefaultServer(gqlgen.NewExecutableSchema(gqlgen.Config{
				Resolvers: gqlResolver,
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
					var gqlErr *gqlerror.Error
					if errors.As(err, &gqlErr) {
						if code, ok := gqlErr.Extensions["code"]; ok && code != errcode.InternalServerError {
							return gqlErr
						}
					}

					slog.Error(err.Error(), slog.Any("error", oops.Wrap(err)))

					return &gqlerror.Error{
						Err:     err,
						Message: "something went wrong",
						Path:    graphql.GetPath(ctx),
						Extensions: map[string]any{
							"code": errcode.InternalServerError,
						},
					}
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
