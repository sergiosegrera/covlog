package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/kevinburke/twilio-go"
	"github.com/sergiosegrera/covlog/models"
	"github.com/sergiosegrera/covlog/service"
)

func MakePostPersonHandler(svc service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request twilio.Message
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			JSON(w, 400, message{"error": "Invalid request, failed to decode message"})
			return
		}

		var person models.Person
		// TODO: Verify name
		person.Name = request.Body
		person.Phone = request.From.Friendly()

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
