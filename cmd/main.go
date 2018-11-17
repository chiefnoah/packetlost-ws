package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"git.packetlostandfound.us/chiefnoah/packetlostandfound-ws/pkg/contact"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func main() {

	POSTGRES_URI := os.Getenv("POSTGRES_URI")

	db, err := sql.Open("postgres", POSTGRES_URI)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	contactService := contact.Service{DB: db}
	contactService.Init()
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		contact.BindEndpoints(r, contactService)
	})

	//Uncomment this to enable static file serving
	// r.Handle("/*", http.FileServer(http.Dir("./static/")))

	//Finally bind everything to the root endpoint
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
