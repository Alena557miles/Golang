package controllers

import (
	"creator/app/models"
	"creator/app/views"
	"fmt"
)

type GalleryController struct {
	Galleries []*models.Gallery
}

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

func (gc *GalleryController) GalleryCreation() {
	// CREATE A GALLERY
	fmt.Printf("Gallery name:\n")
	var name string
	_, _ = fmt.Scan(&name)
	fmt.Printf("Gallery location (city): \n")
	var l string
	_, _ = fmt.Scan(&l)
	gallery := &models.Gallery{Name: name, Location: l}
	// create an gallery and print information about it
	gc.CreateGallery(gallery)
	if err := views.PrintGallery(gallery); err != nil {
		fmt.Println(err)
	}
	fmt.Println("---Gallery created successfuly---")
	fmt.Println("")
}

func (gc *GalleryController) RemoveArtistFromGal(sl Getter) {
	//finding the gallery
	fmt.Printf("From which Gallery ? (Print name of the Gallery):\n")
	var galleryName string
	_, _ = fmt.Scan(&galleryName)
	gallery := gc.FindGallery(galleryName)
	// finding the artist
	fmt.Printf("Which Artist do you want to remove from the Gallery? (Print name of the artist):\n")
	var artistName string
	_, _ = fmt.Scan(&artistName)
	var ar ArsistFinder
	sl.Get(&ar)
	artist := ar.FindArtist(artistName)
	//delete
	gc.DeleteArtist(gallery, artist)
	if err := views.PrintGallery(gallery); err != nil {
		fmt.Println(err)
	}
	fmt.Println("---Artist was removed from the Gallery successfuly---")
	fmt.Println("")
}

type ArsistFinder interface {
	FindArtist(name string) *models.Artist
}
