package entities

import (
)

//Account type is an 'object' for holding accounts
type Account struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Picture   string `json:"picture"`
	IsActive  bool   `json:"is_active"`
	DateCreated string `json:"date_created"`
}



