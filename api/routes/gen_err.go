package routes

import (
	"time"

	"github.com/betelgeusexru/golang-elk/pkg/logging"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func HandleGenError(c *fiber.Ctx) error {
	logger := logging.GetLogger()

	logger.WithFields(logrus.Fields{
		"timestamp": time.Now().UnixNano(),
	}).Info("Hit endpoint HandleGenError")

	resp := map[string]string{"msg": "error"}

	logger.WithFields(logrus.Fields{
		"timestamp": time.Now().UnixNano(),
	}).Error("HandleGenError endpoint error occurred: deliberate error")

	return c.Status(500).JSON(&resp)
}
