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

	// POSTGRES_USER := os.Getenv("POSTGRES_USER")
	// POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	// POSTGRES_ADDR := os.Getenv("POSTGRES_ADDR")

	POSTGRES_URI := os.Getenv("POSTGRES_URI")

	// connStr := fmt.Sprintf("host=%s user=%s password='%s' dbname=packetlostandfound sslmode=disable", POSTGRES_ADDR, POSTGRES_USER, POSTGRES_PASSWORD)
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

	//Finally bind everything to the root endpoint
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
