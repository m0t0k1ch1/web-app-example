package core

import (
	"net/http"

	graphqlplayground "github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"

	"app/config"
)

type Server struct {
	http.Server

	config config.ServerConfig
}

func NewServer(
	srvConf config.ServerConfig,
) *Server {
	var gqlHandler http.Handler
	{
		r := chi.NewRouter()
		{
			// TODO
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
