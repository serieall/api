package controllers

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/prometheus/common/log"
	"github.com/youkoulayley/serieall-api-go/api/bootstrap"
	"github.com/youkoulayley/serieall-api-go/api/models"
	"io/ioutil"
	"net/http"
)

func PublishImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var natsImage models.NatsImage

	err = json.Unmarshal(body, &natsImage)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sc := getStan()

	err = sc.Publish("worker_images", []byte(body))
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(natsImage)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func getStan() stan.Conn {
	return bootstrap.GetStan()
}
