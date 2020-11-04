package models

type Person struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	TTL   int    `json:"ttl"`
}
