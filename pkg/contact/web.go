package contact

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

//BindEndpoints binds service functions to their corresponding RESTful endpoints
func BindEndpoints(r chi.Router, s Service) {
	r.Post("/contact_submission", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received contact submission request")
	})
}
