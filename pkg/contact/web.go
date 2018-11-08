package contact

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func BindEndpoints(r chi.Router, c ContactService) {
	r.Post("/contact_submission", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received contact submission request")
	})
}
