package main

import (
	"github.com/youkoulayley/serieall-api-go/api/controllers"
	"github.com/youkoulayley/serieall-api-go/api/models"
)

var routes = models.Routes{
	models.Route{
		"health",
		"GET",
		"/health",
		controllers.GetHealth,
	},
}
