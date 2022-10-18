package main

import (
	"context"
	"gqlgen/graph"
	"gqlgen/graph/generated"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

func Middleware() func(http.Handler) http.Handler {
	// middleware that adds the x-api-key into context
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			xapikey := r.Header.Get("x-api-key")
			ctx := context.WithValue(r.Context(), "x-api-key", xapikey)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

const defaultPort = "8080"

func main() {

	router := chi.NewRouter()

	router.Use(Middleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
