package logging

import (
	"net"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger = nil

func getLogstashConnection(log *logrus.Logger) net.Conn {
	conn, err := net.Dial("tcp", "localhost:5400")
	if err != nil {
		log.Fatalf("cannot get logstash connection: %v", err)
	}

	return conn
}

func GetLogger() *logrus.Logger {
	return logger
}

func InitLogger() {
	logger = logrus.New()
	logstashConn := getLogstashConnection(logger)
	hook := logrustash.New(logstashConn, logrustash.DefaultFormatter(logrus.Fields{}))
	logger.Hooks.Add(hook)
}
