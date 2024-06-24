package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello Logging")
	logger.WithField("username", "bayu").Info("Hello Logging")

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "sugiantoro")
	entry.Info("Hello Entry")
}