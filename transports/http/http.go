package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sergiosegrera/covlog/service"
	"net/http"
)

func Serve(svc service.Service) error {
	router := chi.NewRouter()
	router.Use(middleware.Compress(5, "gzip"))

	router.Post("/person")
	router.Get("/persons")
	router.Post("/message")

	router.Get("/admin")

	return http.ListenAndServe(":8080", router)
}
