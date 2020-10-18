package handlers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type message map[string]interface{}

type Response struct {
	Message string
}

func JSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func TWIML(w http.ResponseWriter, code int, text string) {
	response := Response{text}
	x, err := xml.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	w.WriteHeader(code)
	w.Write(x)

}
