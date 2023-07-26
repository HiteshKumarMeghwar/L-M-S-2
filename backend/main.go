package main

import (
	"github.com/HiteshKumarMeghwar/L-M-S-2/database"
	"github.com/HiteshKumarMeghwar/L-M-S-2/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	/* Requiring Database Env Variables */
	database.LoadEnvVariables()
	database.Connect()
	app := fiber.New()
	// this is really important ....
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	routes.Setup(app)

	app.Listen(":8080")
}
