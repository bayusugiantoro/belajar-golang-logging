package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "sugiantoro").Info("Hello World")

	logger.WithField("username", "bayu").
		WithField("name", "Bayu Sugiantoro").
		Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "bayu",
		"name":     "Bayu Sugiantoro",
	}).Info("Hello World")
}