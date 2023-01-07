package main

import (
	"net/http"
	"os"
	"pruebas/directives"
	"pruebas/graph"
	"pruebas/middlewares"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(middlewares.AuthMiddleware)

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://192.168.60.179:8000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	authc := graph.Config{Resolvers: &graph.Resolver{}}
	authc.Directives.Auth = directives.Auth

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(authc))

	router.Handle("/", playground.Handler("Starwars", "/query"))
	router.Handle("/query", srv)

	print("localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

}
