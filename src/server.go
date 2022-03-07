package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/database/seeds"
	"github.com/samithiwat/samithiwat-backend/src/route"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/samithiwat/samithiwat-backend/src/config"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/generated"
	graph "github.com/samithiwat/samithiwat-backend/src/graph/resolver"
)

func gqlHandler(resolver *graph.Resolver) http.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	return func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r)
	}
}

func playgroundHandler() http.Handler {
	return playground.Handler("GraphQL playground", "/graphql")
}

func main() {
	loadConfig, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot to load loadConfig", err)
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

	r := route.NewFiberRouter()

	resolver, err := InitializeResolver(client)
	if err != nil {
		fmt.Printf("failed to inject resolver: %s\n", err)
		os.Exit(2)
	}

	r.All("graphql", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(gqlHandler(resolver))(c.Context())
		return nil
	})

	r.All("/", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(playgroundHandler())(c.Context())
		return nil
	})

	if err := r.Listen(":" + strconv.Itoa(loadConfig.App.Port)); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
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
