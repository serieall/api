package main

import (
	"github.com/sirupsen/logrus"
	"github.com/youkoulayley/serieall-api-go/api/bootstrap"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	// Get the version number
	file, err := ioutil.ReadFile("VERSION")
	if err != nil {
		log.Fatal(err)
	}

	v := string(file)
	logrus.Info("Série-All API is loading in v", v)

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
