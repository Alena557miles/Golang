package controllers

import (
	"creator/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ArtistController struct {
	artists []*models.Artist
}

//func (ac *ArtistController) Name() string {
//	return `ArtistController`
//}
//
//func (ac *ArtistController) Path() string {
//	return `/artist/create`
//}
//
//func (ac *ArtistController) DoAction(a string) {
//	switch a {
//	case `Registration`:
//		ac.Registration()
//	}
//}

func (ac *ArtistController) CreateArtist(a *models.Artist) {
	ac.artists = append(ac.artists, a)
}
func (ac *ArtistController) FindArtist(name string) *models.Artist {
	for _, artist := range ac.artists {
		if artist.Name == name {
			return artist
		}
	}
	return nil
}

func (ac *ArtistController) Registration(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var artistName string = vars["artist"]
	//resp := make(map[string]string)
	//resp["message"] = `Artist ` + artistName + ` created successfully`
	artist := &models.Artist{Name: artistName, OnGallery: false}
	ac.CreateArtist(artist)
	jsonResp, err := json.Marshal(artist)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	rw.Write(jsonResp)
}

func (ac *ArtistController) ArtistRegistration(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var artistName string = vars["artist"]
	var galleryName string = vars["gallery"]
	artist := ac.FindArtist(artistName)
	if err := ac.FindArtist(artistName); err != nil {
		galleryC := &GalleryController{}
		gallery := galleryC.FindGallery(galleryName)
		galleryC.RegisterArtist(gallery, artist)
	}
	resp := make(map[string]string)
	resp["message"] = `Artist: ` + artistName + `is registered on Gallery:` + galleryName
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	rw.Write(jsonResp)
}

//type GalleryFinder interface {
//	FindGallery(name string) *models.Gallery
//	RegisterArtist(gallery *models.Gallery, artist *models.Artist)
//}
//
//type MyGetter interface {
//	Get(some any) bool
//}
