package entities

import (
	"fmt"
)

//Account type is an 'object' for holding accounts
type Account struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Picture   string `json:"picture"`
	IsActive  bool   `json:"is_active"`
}

//Getname returns name of object
func (a Account) Getname() string {
	return a.FirstName
}

//Create a new Account struct type
func Create(first string, last string, phone string, picture string, active bool) Account {
	a := Account{FirstName: first, LastName: last, Phone: phone, Picture: picture, IsActive: active}
	return a

}

//Print prints object
func (a Account) Print() {
	fmt.Println(a.FirstName + "\n" + a.LastName + "\n" + a.Phone + "\n" + a.Picture + "\n")
}
