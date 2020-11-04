package handlers

import (
	"context"
	"net/http"

	"github.com/sergiosegrera/covlog/models"
	"github.com/sergiosegrera/covlog/service"
)

func MakePostPersonHandler(svc service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			TWIML(w, 400, "Failed to parse form")
			return
		}

		var person models.Person
		// TODO: Verify name
		person.Name = r.Form.Get("Body")
		person.Phone = r.Form.Get("From")

		// TODO: Context timeout
		err = svc.CreatePerson(context.Background(), person)

		// TODO: Better errors
		if err != nil {
			switch err {
			case service.ErrDatabase:
				TWIML(w, 500, "Database error!")
				return
			default:
				TWIML(w, 500, "Unknown error")
				return
			}
		}

		TWIML(w, 200, "Perfect! Your phone number will be deleted in 14 days")
	}
}
