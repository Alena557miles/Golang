package webservergo

import (
	"context"
	"creator/controllers"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//type ControllersSet interface {
//	Controllers() []Controller
//}
//
//type Controller interface {
//	Path() string
//	Name() string
//	DoAction(string)
//}

func StartServer() {
	ctx := context.Background()
	router := mux.NewRouter()
	srv := &http.Server{
		Addr:              `0.0.0.0:8080`,
		ReadTimeout:       time.Millisecond * 200,
		WriteTimeout:      time.Millisecond * 200,
		IdleTimeout:       time.Second * 10,
		ReadHeaderTimeout: time.Millisecond * 200,
		Handler:           router,
	}
	artistC := &controllers.ArtistController{}
	artC := &controllers.ArtController{}
	galleryC := &controllers.GalleryController{}

	//for _, c := range controllerSet.Controllers() {
	//	switch c.Name() {
	//	case `artC`:
	//		if c.Path() == "/createart/{art}" {
	//			router.HandleFunc(c.Path(), artC.ArtCreation)
	//		} else {
	//			router.HandleFunc(c.Path(), artC.AssignArt)
	//		}
	//
	//	case `artistC`:
	//		if c.Path() == "/createartist/{artist}" {
	//			router.HandleFunc(c.Path(), artistC.Registration)
	//		} else {
	//			router.HandleFunc(c.Path(), artistC.ArtistRegistration)
	//		}
	//
	//	case `galleryC`:
	//		if c.Path() == "/creategallery/{gallery}" {
	//			router.HandleFunc(c.Path(), galleryC.GalleryCreation)
	//		} else {
	//			router.HandleFunc(c.Path(), galleryC.RemoveArtistFromGal)
	//		}
	//	}
	//}
	// CREATE AN ART
	// localhost:8080/createart/blackCat
	router.HandleFunc("/createart/{art}", artC.ArtCreation)

	// CREATE AN ARTIST
	// localhost:8080/createartist/Fillip
	router.HandleFunc("/createartist/{artist}", artistC.Registration)

	// CREATE GALLERY
	// localhost:8080/creategallery/Tokio
	router.HandleFunc("/creategallery/{gallery}", galleryC.GalleryCreation)

	//ASSIGN AN ART TO THE ARTIST (BY NAME)
	// localhost:8080/artist/assign/Fillip/blackCat
	router.HandleFunc("/artist/assign/{artist}/{art}", artC.AssignArt)

	//REGISTRATION AN ARTIST ON THE GALLERY
	// localhost:8080/artist/register/Fillip/Tokio
	router.HandleFunc("/artist/register/{artist}/{gallery}", artistC.ArtistRegistration)

	// DELETE AN ARTIST FROM GALLERY
	// localhost:8080/artist/delete/Fillip/Tokio
	router.HandleFunc("/artist/delete/{artist}/{gallery}", galleryC.RemoveArtistFromGal)

	go func() {
		log.Println(`Web Server started`)
		err := srv.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done,
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	<-done

	log.Println(`Web Server is shutting down`)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(ctx, err)
	}
}
