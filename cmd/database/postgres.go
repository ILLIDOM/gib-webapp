package database

import (
	"database/sql"

	"github.com/ILLIDOM/gin-webapp/cmd/domain"
	_ "github.com/mattn/go-sqlite3"
)

type UserService struct {
	DB *sql.DB
}

func (s *UserService) GetByID(id int) (*domain.User, error) {
	var user domain.User
	row := s.DB.QueryRow(`SELECT id, firstname, lastname FROM users WHERE id = $1`, id)
	err := row.Scan(&user.ID, &user.Firstname, &user.Lastname)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) Create(user domain.User) (*domain.User, error) {
	_, err := s.DB.Exec("INSERT INTO users VALUES(NULL, ?, ?, ?, ?);", user.Firstname, user.Lastname, user.Fullname, user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
