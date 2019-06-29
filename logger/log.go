// Created by Hisen at 2019-06-26.
package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.Formatter = new(logrus.JSONFormatter)
	Logger.Level = logrus.DebugLevel
	Logger.Out = os.Stdout
}
