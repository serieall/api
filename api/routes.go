package main

import (
	"github.com/youkoulayley/serieall-api-go/api/controllers"
	"github.com/youkoulayley/serieall-api-go/api/models"
)

var routes = models.Routes{
	// HEALTH
	models.Route{
		"health",
		"GET",
		"/health",
		controllers.GetHealth,
	},
	// IMAGES
	models.Route{
		"images_url",
		"GET",
		"/images",
		controllers.GetImage,
	},
	models.Route{
		"images",
		"POST",
		"/images",
		controllers.PublishImage,
	},
}
