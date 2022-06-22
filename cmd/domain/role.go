package domain

type Role struct {
	ID          int    `json:"id"`
	Name        string `json:"role_name"`
	Descirption string `json:"role_description"`
}

type RoleService interface {
	GetByID(id int) (*User, error)
	GetAll() ([]*User, error)
	Create(u *User) error
	Delete(u *User) error
}
