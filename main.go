package main

import (
	"github.com/1rvyn/graphql-service/graph"
	"github.com/1rvyn/graphql-service/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	app.All("/query", func(c *fiber.Ctx) error {
		srv.ServeHTTP(c.Context())
		return nil
	})

	app.Listen(":8080")
}
