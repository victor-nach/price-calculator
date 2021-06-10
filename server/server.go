package server

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/victor-nach/price-calculator/db"
	"github.com/victor-nach/price-calculator/graph"
	"github.com/victor-nach/price-calculator/graph/generated"
	"github.com/victor-nach/price-calculator/lib"
	"github.com/victor-nach/price-calculator/lib/rerrors"
	"log"
	"net/http"
	"strings"
)

//Server ...
type Server struct {
	server *handler.Server
}

//NewServer returns a new server
func NewServer(priceService lib.PriceService, dataStore db.Datastore) *Server {
	resolvers := graph.NewResolver(priceService, dataStore)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))

	// set default error presenter
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)

		errString := err.Error()
		err.Message = errString

		index := strings.Index(errString, "{")
		if index != -1 {
			errString = strings.TrimSpace(errString[index:])
		}
		r, er := rerrors.NewErrFromJSON(errString)
		fmt.Println(r, er)
		err.Message = err.Error()
		if er == nil {
			err.Message = r.Message
			err.Extensions = map[string]interface{}{
				"code":      r.Code,
				"errorType": r.ErrorType,
			}
		}

		//err.Message = err.Error()
		//r, ok := e.(*rerrors.Err)
		//fmt.Println(r, ok)
		//fmt.Println()
		//switch v := e.(type) {
		//case *rerrors.Err:
		//	err.Message = v.Message
		//	err.Extensions = map[string]interface{}{
		//		"code":      v.Code,
		//		"errorType": v.ErrorType,
		//	}
		//default:
		//	err.Message = err.Error()
		//}
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
