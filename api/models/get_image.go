package models

type GetImage struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Size string `json:"size"`
}
