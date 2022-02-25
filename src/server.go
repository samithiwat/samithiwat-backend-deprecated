package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/samithiwat/samithiwat-backend/src/config"
	"github.com/samithiwat/samithiwat-backend/src/database"
)

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

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    app.Listen(":" + config.Port)
}