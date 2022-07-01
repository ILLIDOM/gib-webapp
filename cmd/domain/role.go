package domain

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"role_name"`
}

type RoleService interface {
	GetByID(id int) (*User, error)
	GetAll() ([]*User, error)
	Create(u *User) error
	Delete(u *User) error
}
