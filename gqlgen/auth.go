package auth

import (
	"context"
	"database/sql"
	"net/http"
)

func Middleware(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Header

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

//if there is a request to our graphql server we want to look at the headers

//pass the context with the header

//middleware that checks the headers of the request in x-api-key and assigns them to context
// request headers
// send them to context
// with the request
