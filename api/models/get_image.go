package models

type GetImage struct {
	Id   int    `json:"id"`
	Url  string `json:"url"`
	Name string `json:"name"`
	Type string `json:"type"`
	Size string `json:"size"`
}
