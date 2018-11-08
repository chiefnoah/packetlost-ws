package contact

import (
	"database/sql"
	"log"
	"time"
)

type ContactSubmission struct {
	ID      uint64    `json:"id"`
	Email   string    `json:"email"`
	Message string    `json:"message"`
	Date    time.Time `json:"submission_date"`
}

type ContactService struct {
	DB *sql.DB
}

func (c *ContactService) Init() {
	if _, err := c.DB.Exec("CREATE TABLE IF NOT EXISTS CONTACT_SUBMISSIONS(ID INT PRIMARY KEY, EMAIL TEXT NOT NULL, MESSAGE TEXT NOT NULL, SUBMITTED DATE NOT NULL DEFAULT CURRENT_DATE)"); err != nil {
		log.Fatal(err)
	}
}
