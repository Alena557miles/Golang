package controllers

import (
	"creator/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"fmt"
)

type ArtController struct {
	arts []*models.Art
}

//func (ac *ArtController) Name() string {
//	return `ArtController`
//}
//
//func (ac *ArtController) Path() string {
//	return `art?action=create`
//}
//
//func (ac *ArtController) DoAction(a string) {
//	switch a {
//	case `ArtCreation`:
//	}
//}

func (ac *ArtController) CreateArt(a *models.Art) {
	ac.arts = append(ac.arts, a)
}

func (ac *ArtController) FindArt(name string) *models.Art {
	for _, art := range ac.arts {
		if art.Name == name {
			return art
		}
	}
	return nil
}

func (ac *ArtController) AssignedArtToArtist(art *models.Art, artist *models.Artist) *models.Art {
	if art.IsntAssigned() {
		art.Owner = artist.Name
		artist.Arts = append(artist.Arts, art)
		fmt.Println(artist.Name)
		return art
	} else {
		fmt.Println("This art already has an owner! You'd better make your own art")
	}
	return nil
}

func (ac *ArtController) ArtCreation(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var artName string = vars["art"]
	//resp := make(map[string]string)
	//resp["message"] = `Art ` + artName + ` created successfully`
	art := &models.Art{Name: artName}
	ac.CreateArt(art)
	jsonResp, err := json.Marshal(art)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	rw.Write(jsonResp)
}

func (ac *ArtController) AssignArt(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var artistName string = vars["artist"]
	var artName string = vars["art"]
	art := ac.FindArt(artName)
	if err := ac.FindArt(artName); err != nil {
		artistC := &ArtistController{}
		artist := artistC.FindArtist(artistName)
		if err := artistC.FindArtist(artistName); err != nil {
			//assign an art to the Artist and
			ac.AssignedArtToArtist(art, artist)
		}
	}
	resp := make(map[string]string)
	resp["message"] = `Art: ` + artName + ` is assigned to Artist:` + artistName
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	rw.Write(jsonResp)
}

//type ArtistFinder interface {
//	FindArtist(name string) *models.Artist
//}
//
//type Getter interface {
//	Get(some any) bool
//}
