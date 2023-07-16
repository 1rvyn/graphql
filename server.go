package main

import (
	"log"
	"net/http"
	"os"

	"github.com/1rvyn/graphql-service/database"
	"github.com/1rvyn/graphql-service/graph"
	"github.com/1rvyn/graphql-service/routes"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	database.ConnectDb()

	router := mux.NewRouter()

	setUpRoutes(router)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/employee", srv)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func setUpRoutes(router *mux.Router) {
	router.Handle("/login", routes.Login(router))
}
