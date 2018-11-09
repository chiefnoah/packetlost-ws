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
	if _, err := c.DB.Exec("CREATE DATABASE IF NOT EXISTS packetlostandfound"); err != nil {
		log.Fatal(err)
	}
	if _, err := c.DB.Exec("CREATE TABLE IF NOT EXISTS CONTACT_SUBMISSIONS(ID SERIAL PRIMARY KEY, EMAIL TEXT NOT NULL, MESSAGE TEXT NOT NULL, SUBMITTED DATE NOT NULL DEFAULT CURRENT_DATE)"); err != nil {
		log.Fatal(err)
	}
}

//Create creates and stores a contact submission in the database
func (c *Service) Create(cs *Submission) error {
	_, err := c.DB.Exec("INSERT INTO CONTACT_SUBMISSIONS(EMAIL, MESSAGE) VALUES($1, $2)", cs.Email, cs.Message)
	return err
}

func (c *Service) List() (*[]Submission, error) {
	rows, err := c.DB.Query("SELECT ID, EMAIL, MESSAGE, SUBMITTED FROM CONTACT_SUBMISSIONS")
	if err != nil {
		log.Printf("Unable to get submissions: %s", err)
		return nil, err
	}
	submissions := make([]Submission, 0, 5)
	for rows.Next() {
		var sub Submission
		if err := rows.Scan(&sub.ID, &sub.Email, &sub.Message, &sub.Date); err != nil {
			log.Printf(err.Error())
		}
		submissions = append(submissions, sub)
	}
	return &submissions, nil
}
