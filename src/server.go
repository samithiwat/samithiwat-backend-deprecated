package main

import (
	"flag"
	"fmt"
	"github.com/samithiwat/samithiwat-backend/src/database/seeds"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"

	"github.com/samithiwat/samithiwat-backend/src/config"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/generated"
	graph "github.com/samithiwat/samithiwat-backend/src/graph/resolver"

	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func gqlHandler(resolver *graph.Resolver) http.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r)
	})
}

func playgroundHandler() http.Handler {
	return playground.Handler("GraphQL playground", "/graphql")
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

	handleArgs(client)

	app := fiber.New()

	resolver, err := InitializeResolver(client)
	if err != nil {
		fmt.Printf("failed to inject resolver: %s\n", err)
		os.Exit(2)
	}

	app.All("graphql", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(gqlHandler(resolver))(c.Context())
		return nil
	})

	app.All("/", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(playgroundHandler())(c.Context())
		return nil
	})

	app.Listen(":" + strconv.Itoa(config.App.Port))
}

func handleArgs(db database.Database) {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			err := seed.Execute(db, args[1:]...)
			if err != nil {
				log.Fatalln("Not found seed")
			}
			os.Exit(0)
		}
	}
}
