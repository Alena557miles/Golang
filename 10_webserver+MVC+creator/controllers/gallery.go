package controllers

import (
	"creator/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type GalleryController struct {
	Galleries []*models.Gallery
}

//func (ac *GalleryController) Name() string {
//	return `GalleryController`
//}
//
//func (ac *GalleryController) Path() string {
//	return `/gallery`
//}
//
//func (ac *GalleryController) DoAction(a string) {
//	switch a {
//	case `GalleryCreation`:
//		ac.GalleryCreation()
//	}
//}

func (gc *GalleryController) CreateGallery(g *models.Gallery) {
	gc.Galleries = append(gc.Galleries, g)
}
func (gc *GalleryController) FindGallery(name string) *models.Gallery {
	for _, g := range gc.Galleries {
		if name == g.Name {
			return g
		}
	}
	return nil
}

func (gc *GalleryController) RegisterArtist(gallery *models.Gallery, artist *models.Artist) {

	if len(artist.Arts) > 0 {
		gc.Galleries = append(gc.Galleries, gallery)
		gallery.Artists = append(gallery.Artists, artist)
		artist.OnGallery = true
		return
	}
	if len(artist.Arts) == 0 {
		fmt.Println("We can not register an Artist without Arts")
	}
	if artist.OnGallery {
		fmt.Println("Artist are already on gallery")
	}
}
func (gc *GalleryController) DeleteArtist(gallery *models.Gallery, artist *models.Artist) {
	for _, g := range gc.Galleries {
		if g == gallery {
			gallery.DeleteArtist(artist)
		}
	}
}

func (gc *GalleryController) GalleryCreation(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var galleryName string = vars["gallery"]
	resp := make(map[string]string)
	resp["message"] = `Gallery ` + galleryName + ` created successfully`
	gallery := &models.Gallery{Name: galleryName}
	gc.CreateGallery(gallery)
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	rw.Write(jsonResp)
}

func (gc *GalleryController) RemoveArtistFromGal(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var artistName string = vars["artist"]
	var galleryName string = vars["gallery"]
	artistC := &ArtistController{}
	artist := artistC.FindArtist(artistName)
	if err := artistC.FindArtist(artistName); err != nil {
		gallery := gc.FindGallery(galleryName)
		gc.DeleteArtist(gallery, artist)
	}
	resp := make(map[string]string)
	resp["message"] = `Artist:` + artistName + `is deleted from Gallery:` + galleryName
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	rw.Write(jsonResp)
}

//type ArsistFinder interface {
//	FindArtist(name string) *models.Artist
//}
