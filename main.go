package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/kriangkrai/Deivces/Controller"
	"github.com/kriangkrai/Deivces/Router"
)

func main() {

	if err := Controller.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	api := app.Group("/api") // /api
	api.Get("/get/:device", Router.Get)
	api.Get("/gets", Router.Gets)
	api.Post("/insert", Router.Insert)
	api.Put("/update", Router.Update)
	api.Delete("/delete/:device", Router.Delete)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app.Listen(":" + port)
	//log.Fatal(app.Listen(port))
}
