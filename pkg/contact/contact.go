package contact

import (
	"database/sql"
	"log"
	"time"
)

//Submission defines the structure for a contact request
type Submission struct {
	ID      uint64    `json:"id"`
	Email   string    `json:"email"`
	Message string    `json:"message"`
	Date    time.Time `json:"submission_date"`
}

//Service defines a service for handling submitted contact forms
type Service struct {
	DB *sql.DB
}

//Init initializes the service and creates any tables necessary in Service.DB to run the service
func (c *Service) Init() {
	if _, err := c.DB.Exec("CREATE TABLE IF NOT EXISTS CONTACT_SUBMISSIONS(ID INT PRIMARY KEY, EMAIL TEXT NOT NULL, MESSAGE TEXT NOT NULL, SUBMITTED DATE NOT NULL DEFAULT CURRENT_DATE)"); err != nil {
		log.Fatal(err)
	}
}

//Create creates and stores a contact submission in the database
func (c *Service) Create(cs *Submission) {
	panic("Not implemented")
}
