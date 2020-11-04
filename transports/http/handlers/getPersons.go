package handlers

import (
	"context"
	"net/http"

	"github.com/sergiosegrera/covlog/service"
)

func MakeGetPersonsHandler(svc service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Context timeout
		phones, err := svc.GetPersons(context.Background())

		// TODO: Better errors
		if err != nil {
			switch err {
			case service.ErrDatabase:
				JSON(w, 500, message{"error": "Database error!"})
				return
			default:
				JSON(w, 500, message{"Unknown error": "Unknown error!"})
				return
			}
		}

		JSON(w, 200, message{"message": phones})
	}
}
