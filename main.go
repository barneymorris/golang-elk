package main

import (
	"time"

	"github.com/betelgeusexru/golang-elk/api/routes"
	"github.com/betelgeusexru/golang-elk/pkg/logging"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	logging.InitLogger()
	log := logging.GetLogger()

	routes := routes.GetRoutes()

	app := fiber.New()

	app.Get("/api/hello", routes.HandleHello)
	app.Get("/api/error", routes.HandleGenError)

	log.WithFields(logrus.Fields{
		"timestamp": time.Now().UnixNano(),
	}).Info("Starting golang-elk server on port 3000")

	err := app.Listen(":3000")
	if err != nil {
		log.WithFields(logrus.Fields{
			"timestamp": time.Now().UnixNano(),
		}).Fatalf("Failed to start golang-elk server on port 3000")
	}
}
