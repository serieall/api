package bootstrap

import (
	"github.com/caarlos0/env"
	"github.com/serieall/api/api/models"
	"github.com/sirupsen/logrus"
)

var Config models.Config

func InitConfig() models.Config {
	err := env.Parse(&Config)
	if err != nil {
		logrus.Error(err)
	}

	return Config
}

func GetConfig() models.Config {
	return Config
}
