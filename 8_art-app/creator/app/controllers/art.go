package controllers

import (
	"creator/app/models"
	"creator/app/views"
	"fmt"
)

type ArtController struct {
	arts []*models.Art
}

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

func (ac *ArtController) ArtCreation() {
	// CREATE AN ART
	fmt.Printf("Art name:\n")
	var name string
	_, _ = fmt.Scan(&name)
	fmt.Printf("Width, cm (only numbers): \n")
	var w int
	_, _ = fmt.Scan(&w)
	fmt.Printf("Height, cm (only numbers): \n")
	var h int
	_, _ = fmt.Scan(&h)
	art := &models.Art{Name: name, Width: w, Height: h}
	// create an art and print information about art
	ac.CreateArt(art)
	if err := views.PrintArt(art); err != nil {
		fmt.Println(err)
	}
	fmt.Println("---________Art created successfuly________---")
	fmt.Println("")
}

func (ac *ArtController) AssignArt(sl Getter) {
	//ASSIGN AN ART TO THE ARTIST (BY NAME)
	//asking and searching an art by name
	fmt.Printf("Which art do you want to assign? (Print name of the art):\n")
	var name string
	_, _ = fmt.Scan(&name)
	art := ac.FindArt(name)
	if err := ac.FindArt(name); err != nil {
		//asking and searching an artist by name
		fmt.Printf("To which artist do you want to assign this art? (Print name of the artist):\n")
		var artistName string
		_, _ = fmt.Scan(&artistName)

		var artistC ArtistFinder
		sl.Get(&artistC)
		artist := artistC.FindArtist(artistName)
		if err := artistC.FindArtist(artistName); err != nil {
			//assign an art to the Artist and
			ac.AssignedArtToArtist(art, artist)
		}
	}
	// print information about art
	if err := views.PrintArt(art); err != nil {
		fmt.Println(err)
	}
	fmt.Println("---Art was assigned to Artist successfuly---")
	fmt.Println("")
}

type ArtistFinder interface {
	FindArtist(name string) *models.Artist
}

type Getter interface {
	Get(some any) bool
}
