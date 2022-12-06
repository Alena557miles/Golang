package models

type Artist struct {
	Name      string `json:"name"`
	Arts      []*Art `json:"art"`
	OnGallery bool   `json:"ongallery"`
}
