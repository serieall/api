package controllers

import (
	"encoding/json"
	"github.com/serieall/api/api/models"
	"net/http"
)

func GetHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.Healthcheck{Status: "OK", Code: 200})
}
