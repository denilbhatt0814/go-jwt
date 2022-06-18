package main

import (
	"github.com/denilbhatt0814/go-jwt/database"
	"github.com/denilbhatt0814/go-jwt/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)


func main(){
	database.Connect()


	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.SetupRoutes(app)
	

	app.Listen(":8585")
}