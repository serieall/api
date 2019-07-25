package controllers

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/prometheus/common/log"
	"github.com/youkoulayley/serieall-api-go/api/bootstrap"
	"github.com/youkoulayley/serieall-api-go/api/models"
	"io/ioutil"
	"net/http"
	"os"
)

func GetImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var getImage models.GetImage

	err = json.Unmarshal(body, &getImage)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var imageUrl models.ImageURL
	var imageFolder = bootstrap.GetConfig().ImageFolder

	imageSizePath := imageFolder + "/" + getImage.Size + "/" + getImage.Name + "_" + getImage.Type + ".jpg"
	imageOriginalPath := imageFolder + "/original/" + getImage.Name + "_" + getImage.Type + ".jpg"
	imageDefaultPath := imageFolder + "/original/default.jpg"

	_, imageSize := os.Stat(imageSizePath)
	if os.IsNotExist(imageSize) {
		_, imageOriginal := os.Stat(imageOriginalPath)
		if os.IsNotExist(imageOriginal) {

			imageUrl.Url = imageDefaultPath
		} else {
			imageUrl.Url = imageOriginalPath
		}
	} else {
		imageUrl.Url = imageSizePath
	}

	err = json.NewEncoder(w).Encode(imageUrl)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

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
