package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sergiosegrera/covlog/config"
	"github.com/sergiosegrera/covlog/service"
	"github.com/sergiosegrera/covlog/transports/http/handlers"
)

func Serve(svc service.Service, conf *config.Config) error {
	router := chi.NewRouter()
	router.Use(middleware.Compress(5, "gzip"))
	router.Use(middleware.Logger)

	router.Post("/person", handlers.MakePostPersonHandler(svc))
	router.Get("/persons", handlers.MakeGetPersonsHandler(svc))

	// TODO: Create csv :o
	// router.Get("/download")

	// router.Get("/admin")
	// router.Get("/frontend")

	return http.ListenAndServe(fmt.Sprintf(":%v", conf.HttpPort), router)
}
