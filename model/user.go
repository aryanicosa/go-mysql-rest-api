package model

type User struct {
	ID			int		`json:"id"`
	Name		string	`json:"name"`
	Email		string	`Json:"email" sql:"unique"`
	Address		string	`json:"address"`
	Password	[]byte	`json:"-"`
}