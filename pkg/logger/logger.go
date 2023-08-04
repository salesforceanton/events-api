package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func LogHandlerIssue(handler string, err error) {
	logrus.WithFields(logrus.Fields{
		"handler": handler,
		"problem": fmt.Sprintf("Error appeared in handler [%s]: %s", handler, err.Error()),
	}).Error(err)
}

func LogExecutionIssue(err error) {
	logrus.WithFields(logrus.Fields{
		"handler": "main",
		"problem": fmt.Sprintf("Error during application running: %s", err.Error()),
	}).Error(err)
}
