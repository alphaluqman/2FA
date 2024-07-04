package main

import (
	"fmt"
	"github.com/alphaluqman/2FA/intenal/api"
	"github.com/alphaluqman/2FA/intenal/client"
	"github.com/alphaluqman/2FA/intenal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return
	}
	app := fiber.New()

	// initialize twillio clinet
	twilioClient := client.InitTwilioClient()

	// initialize twwillio service and handler
	authService := service.NewAuthService(twilioClient)
	authHandler := api.NewAuthHandler(authService)
	api.SetupRoutes(app, authHandler)

	//go func() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	if err = app.Listen(port); err != nil {
		fmt.Println(err)
		//log.Fatal(err)
	}
	//}()

}
