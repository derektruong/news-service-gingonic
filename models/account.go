package models


type Account struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func NewAccount(id int, name, email, password string) *Account {
	return &Account{
		ID: id,
		Name: name,
		Email: email,
		Password: password,
	}
}