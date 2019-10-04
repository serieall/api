package main

import (
	"github.com/serieall/api/api/controllers"
	"github.com/serieall/api/api/models"
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
