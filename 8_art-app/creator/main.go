package main

import (
	"creator/app/controllers"
	"creator/app/sl"
	"fmt"
)

func main() {

	artistC := &controllers.ArtistController{}
	artC := &controllers.ArtController{}
	galleryC := &controllers.GalleryController{}
	sl := &sl.ServiceLocator{}

START:
	fmt.Printf("APP 'CREATOR' CHOOSE WHAT TO DO:" +
		"\nRegister an Artist: ra" +
		"\nCreate an art: ca" +
		"\nAdd art to the Artist: aa" +
		"\nCreate a Gallery: cg" +
		"\nRegister an Artist in Gallery: rg" +
		"\nDelete an Artist from the Gallery: da" +
		"\nYou want to exit: e\n")

	var command string
	_, _ = fmt.Scan(&command)

	switch command {
	case `ra`:
		// ARTIST REGISTRATION
		artistC.Registration()
		goto START

	case `ca`:
		// CREATE AN ART
		artC.ArtCreation()
		goto START

	case `cg`:
		// CREATE A GALLERY
		galleryC.GalleryCreation()
		goto START

	case `aa`:
		//ASSIGN AN ART TO THE ARTIST (BY NAME)
		sl.Register(artistC)
		artC.AssignArt(sl)
		goto START

	case `rg`:
		//REGISTRATION AN ARTIST ON THE GALLERY
		sl.Register(galleryC)
		artistC.ArtistRegistration(sl)
		goto START

	case `da`:
		// DELETE AN ARTIST FROM GALLERY
		sl.Register(artistC)
		galleryC.RemoveArtistFromGal(sl)
		goto START
	case `e`:
		break
	}
	fmt.Printf("as you wish...bye...")
}
