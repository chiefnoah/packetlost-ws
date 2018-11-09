package contact

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

//BindEndpoints binds service functions to their corresponding RESTful endpoints
func BindEndpoints(r chi.Router, s Service) {
	r.Post("/contact_submission", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s called...", r.URL.Path)
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		var submission Submission
		err := decoder.Decode(&submission)
		if err != nil {
			log.Printf("Called: %s - Invalid JSON payload: %s", r.URL.Path, err)
			http.Error(w, "Invalid JSON request", http.StatusBadRequest)
			return
		}
		err = s.Create(&submission)
		if err != nil {
			log.Printf("Error creating submission: %s", err)
			http.Error(w, "Unable to submit", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
	})
	r.Get("/list_submissions", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s called...", r.URL.Path)
		subs, err := s.List()
		if err != nil {
			log.Printf("Unable to get submissions: %s", err.Error())
			http.Error(w, "Unable to list submissions", http.StatusInternalServerError)
		}
		responseText, err := json.Marshal(subs)
		if err != nil {
			errorResponse := fmt.Sprintf("Unable to marshal JSON response: %s", err)
			log.Printf(errorResponse)
			http.Error(w, errorResponse, http.StatusInternalServerError)
			return
		}
		w.Write(responseText)
	})
}
