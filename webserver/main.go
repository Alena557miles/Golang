package webserver

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type ControllerSet interface {
	Controllers() []Controller
}

type Controller interface {
	CreateArt()
}

func SatrtServer(cs ControllerSet) {
	ctx := context.Background()
	rtr := mux.NewRouter()
	srv := &http.Server{
		Addr:              `0.0.0.0:8080`,
		ReadTimeout:       time.Millisecond * 200,
		WriteTimeout:      time.Millisecond * 200,
		IdleTimeout:       time.Second * 10,
		ReadHeaderTimeout: time.Millisecond * 200,
		Handler:           rtr,
	}

	rtr.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		_, _ = rw.Write([]byte("Hello!"))
	})
	rtr.HandleFunc("/healthcheck", func(rw http.ResponseWriter, r *http.Request) {
		_, _ = rw.Write([]byte(`OK`))
	})
	rtr.HandleFunc("/point3", func(rw http.ResponseWriter, r *http.Request) {
		_, _ = rw.Write([]byte(`I'm point 3'`))
	})

	http.Handle("/", rtr)

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
