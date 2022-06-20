package domain

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
}

type UserService interface {
	GetByID(id int) (*User, error)
	GetAll() ([]*User, error)
	Create(u *User) error
	Delete(u *User) error
}
