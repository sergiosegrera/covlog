package models

import "time"

type Person struct {
	Name    string    `json:"name"`
	Phone   string    `json:"phone"`
	Date    time.Time `json:"date"`
	Expires time.Time `json:"expires"`
}
