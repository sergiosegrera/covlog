package http

import (
	"net/http"
	"fmt"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sergiosegrera/covlog/service"
	"github.com/sergiosegrera/covlog/transports/http/handlers"
)

func Serve(svc service.Service, conf *config.Config) error {
	router := chi.NewRouter()
	router.Use(middleware.Compress(5, "gzip"))

	router.Post("/person", handlers.MakePostPersonHandler(svc))
	// router.Get("/persons")
	// router.Post("/message")

	// router.Get("/admin")
	// router.Get("/frontend")

	return http.ListenAndServe(fmt.Sprintf(":%v", conf.HttpPort), router)
}
