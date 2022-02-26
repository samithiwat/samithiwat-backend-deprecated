package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/config"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/generated"
	graph "github.com/samithiwat/samithiwat-backend/src/graph/resolver"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func gqlHandler(resolver *graph.Resolver) http.HandlerFunc{
    srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        srv.ServeHTTP(w, r)
    })
}

func playgroundHandler() http.HandlerFunc{
    return playground.Handler("GraphQL playground", "/query")
}

func main() {
    config, err := config.LoadConfig(".")
    
    if err != nil {
        log.Fatal("cannot to load config", err)
    }

    client, err := database.InitDatabase()
    
    if err != nil {
        log.Fatal("cannot to init database", err)
    }

    err = client.AutoMigrate()
    
    if err != nil {
        log.Fatal("cannot migrate database", err)
    }

    app := fiber.New()

// app.All("query", func(c *fiber.Ctx) error {
//     fasthttpadaptor.NewFastHTTPHandler(gqlHandler())(c.Context())
// })

    app.All("/", func(c *fiber.Ctx) error {
        fasthttpadaptor.NewFastHTTPHandler(playgroundHandler())(c.Context())
        return nil
    })

    app.Listen(":" + config.Port)
}