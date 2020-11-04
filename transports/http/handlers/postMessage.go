package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sergiosegrera/covlog/service"
)

type PostMessageRequest struct {
	Message string `json:"message"`
}

func MakePostMessageHandler(svc service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request PostMessageRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			JSON(w, 400, message{"error": "Invalid request, failed to decode"})
			return
		}

		err = svc.SendMessages(r.Context(), request.Message)
		if err != nil {
			JSON(w, 500, message{"error": "Failed sending message"})
			return
		}

		JSON(w, 200, message{"message": "Success"})
	}
}
