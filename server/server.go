package server

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/victor-nach/price-calculator/graph"
	"github.com/victor-nach/price-calculator/graph/generated"
	"github.com/victor-nach/price-calculator/lib"
	"github.com/victor-nach/price-calculator/lib/rerrors"
	"log"
	"net/http"
)

//Server ...
type Server struct {
	server *handler.Server
}

//NewServer returns a new server
func NewServer(priceService lib.PriceService) *Server {
	resolvers := graph.NewResolver(priceService)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))

	// set default error presenter
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)

		err.Message = err.Error()
		var myErr *rerrors.Err
		if errors.As(e, &myErr) {
			myErr, ok := e.(*rerrors.Err)
			if ok {
				err.Message = myErr.Message
				err.Extensions = map[string]interface{}{
					"code":      myErr.Code,
					"errorType": myErr.ErrorType,
				}
			}
		}
		return err
	})

	return &Server{server: srv}
}

//Run starts the server on a specified address
func (s *Server) Run(address string) error {
	http.Handle("/", playground.Handler("GraphQL playground", "/graphiql"))
	http.Handle("/graphiql", s.server)

	log.Printf("connect to http://localhost%s/ for GraphQL playground", address)
	return http.ListenAndServe(address, nil)
}
