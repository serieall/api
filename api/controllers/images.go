package controllers

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/youkoulayley/serieall-api-go/api/bootstrap"
	"github.com/youkoulayley/serieall-api-go/api/models"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func GetImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var getImage models.GetImage

	err = json.Unmarshal(body, &getImage)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var imageUrl models.ImageURL
	var imageFolder = bootstrap.GetConfig().ImageFolder
	var natsImage models.NatsImage

	if getImage.Type == "poster" {
		natsImage.Url = "https://www.thetvdb.com/banners/posters/" + strconv.Itoa(getImage.Id) + "-1.jpg"
		natsImage.CropType = "poster"
	} else {
		natsImage.Url = "https://www.thetvdb.com/banners/graphical/" + strconv.Itoa(getImage.Id) + "-g2.jpg"
		natsImage.CropType = "banner"
	}

	imageSizePath := imageFolder + "/" + getImage.Size + "/" + getImage.Name + "-" + getImage.Type + ".jpg"
	imageOriginalPath := imageFolder + "/original/" + getImage.Name + "-" + getImage.Type + ".jpg"
	imageDefaultPath := imageFolder + "/original/default.jpg"

	natsImage.Name = getImage.Name
	natsImage.Crop = "middle"
	natsImage.ForceCrop = true

	_, imageSize := os.Stat(imageSizePath)
	if os.IsNotExist(imageSize) {
		_, imageOriginal := os.Stat(imageOriginalPath)
		if os.IsNotExist(imageOriginal) {
			err = publishImage(natsImage)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			imageUrl.Url = imageDefaultPath
		} else {
			imageUrl.Url = imageOriginalPath
			err = publishImage(natsImage)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	} else {
		imageUrl.Url = imageSizePath
	}

	err = json.NewEncoder(w).Encode(imageUrl)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func PublishImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var natsImage models.NatsImage

	err = json.Unmarshal(body, &natsImage)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sc := getStan()

	err = sc.Publish("worker_images", []byte(body))
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(natsImage)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func publishImage(natsImage models.NatsImage) error {
	sc := getStan()

	body, err := json.Marshal(natsImage)
	if err != nil {
		logrus.Error(err)
		return error(err)
	}

	logrus.Debugf("Sending images %v to worker_images", natsImage)
	err = sc.Publish("worker_images", []byte(body))
	if err != nil {
		logrus.Error(err)
		return error(err)
	}

	return nil
}

func getStan() stan.Conn {
	return bootstrap.GetStan()
}
