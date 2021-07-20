package main

import (
	"github.com/sirupsen/logrus"
)

var logger = logrus.WithFields(logrus.Fields{})

func main() {
	logger.Info("test")
}
