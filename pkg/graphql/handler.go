package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"

	"hatflix/pkg/graphql/gqlgen"
	graphql "hatflix/pkg/graphql/resolvers"
)

func NewHandler(db *sqlx.DB) http.Handler {
	r := chi.NewRouter()
	gqlgenConfig := gqlgen.Config{
		Resolvers: graphql.NewResolverRoot(db),
	}
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	svc := handler.NewDefaultServer(gqlgen.NewExecutableSchema(gqlgenConfig))

	r.Handle("/", svc)
	r.Handle("/explorer", playground.Handler("Explorer", "/"))
	return r
}
