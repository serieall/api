package models

type NatsImage struct {
	Url       string `json:"url"`
	Name      string `json:"name"`
	CropType  string `json:"crop_type"`
	Crop      string `json:"crop"`
	ForceCrop bool   `json:"force_crop"`
}
