package main

import (
	"github.com/serieall/api/api/bootstrap"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	logrus.Info("Série-All API is loading in v0.1")

	// Get the configuration
	cfg := bootstrap.InitConfig()
	bootstrap.InitStan()

	// Setup log level
	var level logrus.Level
	switch strings.ToLower(cfg.LogLevel) {
	case "debug":
		level = logrus.DebugLevel
	case "info":
		level = logrus.InfoLevel
	case "warning":
		level = logrus.WarnLevel
	case "error":
		level = logrus.ErrorLevel
	default:
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)

	logrus.Debugf("Env loaded: %+v", cfg)

	portStr := strconv.Itoa(cfg.Port)

	// Initialize router
	r := initializeRouter() // Init router
	logrus.Infof("Start to listen on %s ...", portStr)
	logrus.Fatal(http.ListenAndServe(":"+portStr, r))
}
