package controllers

import (
	"creator/app/models"
	"creator/app/views"
	"fmt"
)

type ArtistController struct {
	artists []*models.Artist
}

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

func (ac *ArtistController) Registration() {
	// ARTIST REGISTRATION
	fmt.Printf("Artist name:\n")
	var name string
	_, _ = fmt.Scan(&name)
	artist := &models.Artist{Name: name, OnGallery: false}
	// create an Artist and print information about user
	ac.CreateArtist(artist)
	if err := views.PrintArtist(artist); err != nil {
		fmt.Println(err)
	}
	fmt.Println("---Artist registrated successfuly---")
	fmt.Println("")
}

func (ac *ArtistController) ArtistRegistration(sl MyGetter) {
	//asking and searching an artist by name
	fmt.Printf("Which artist do you want to registrate on the gallery? (Print name of the artist):\n")
	var artistName string
	_, _ = fmt.Scan(&artistName)
	artist := ac.FindArtist(artistName)
	if err := ac.FindArtist(artistName); err != nil {
		fmt.Printf("To which Gallery ? (Print name of the Gallery):\n")
		var galleryName string
		_, _ = fmt.Scan(&galleryName)
		//searching gallery
		var gal GalleryFinder
		sl.Get(&gal)
		gallery := gal.FindGallery(galleryName)
		gal.RegisterArtist(gallery, artist)
		// Artist registration on the Gallery and print information about artist
		if err := views.PrintArtist(artist); err != nil {
			fmt.Println(err)
		}
		fmt.Println("---Artist was registrated on the Gallery successfuly---")
	}
	fmt.Println("")
}

type GalleryFinder interface {
	FindGallery(name string) *models.Gallery
	RegisterArtist(gallery *models.Gallery, artist *models.Artist)
}

type MyGetter interface {
	Get(some any) bool
}
