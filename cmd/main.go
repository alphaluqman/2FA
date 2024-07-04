package main

import (
	"github.com/alphaluqman/2FA/intenal/api"
	"github.com/alphaluqman/2FA/intenal/client"
	"github.com/alphaluqman/2FA/intenal/service"
	helper "github.com/alphaluqman/2FA/utils"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	helper.Init()
	app := fiber.New()
	// initialize twillio clinet
	twilioClient := client.InitTwilioClient()

	// initialize twwillio service and handler
	authService := service.NewAuthService(twilioClient)
	authHandler := api.NewAuthHandler(authService)
	api.SetupRoutes(app, authHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	// start server
	if err := app.Listen(port); err != nil {
		log.Fatal(err)
	}

}
