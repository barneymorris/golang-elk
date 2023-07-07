package routes

import (
	"time"

	"github.com/betelgeusexru/golang-elk/pkg/logging"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func HandleHello(c *fiber.Ctx) error {
	logger := logging.GetLogger()

	logger.WithFields(logrus.Fields{
		"timestamp": time.Now().UnixNano(),
	}).Info("Hit endpoint HandleGenError")

	type TResponse struct {
		msg string
	}

	resp := map[string]string{"msg": "hello world"}

	logger.WithFields(logrus.Fields{
		"timestamp": time.Now().UnixNano(),
	}).Infof("HandleGenError endpoint return response: %v", resp)

	return c.JSON(resp)
}
