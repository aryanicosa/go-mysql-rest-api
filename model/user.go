package model

type User struct {
	ID		int		`json:"id"`
	Name	string	`json:"name"`
	Email	string	`Json:"email"`
	Address	string	`json:"address"`
}