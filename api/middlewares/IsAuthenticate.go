package middlewares

import (
	"encoding/json"
	"github.com/youkoulayley/serieall-api-go/api/models"
	"github.com/youkoulayley/serieall-api-go/api/bootstrap"
	"net/http"
	"strings"
)

func IsAuthenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				if bearerToken[1] == bootstrap.GetConfig().Secret {
					next.ServeHTTP(w, r)
				} else {
					json.NewEncoder(w).Encode(models.JSONError{Message: "Invalid authorization token", Code: 403})
				}
			} else {
				json.NewEncoder(w).Encode(models.JSONError{Message: "Invalid authorization token", Code: 403})
			}
		} else {
			json.NewEncoder(w).Encode(models.JSONError{Message: "An authorization header is required", Code: 403})
		}
	})
}