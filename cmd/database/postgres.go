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
	row := s.DB.QueryRow(`SELECT id, firstname, lastname, fullname, email, password FROM users WHERE id = $1`, id)
	err := row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Fullname, &user.Email, &user.Password)
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

func (s *UserService) GetRolesByUserID(userID int) ([]domain.Role, error) {
	userRoles := []domain.Role{}
	rows, err := s.DB.Query(`SELECT roleid FROM userrole WHERE userid = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roleIDs := []int{}
	for rows.Next() {
		var roleID int
		if err := rows.Scan(&roleID); err != nil {
			return nil, err
		}
		roleIDs = append(roleIDs, roleID)
	}

	for _, roleID := range roleIDs {
		role, err := s.GetRoleByID(roleID)
		if err != nil {
			return nil, err
		}
		userRoles = append(userRoles, *role)
	}

	return userRoles, nil
}

func (s *UserService) GetRoleByID(roleID int) (*domain.Role, error) {
	var role domain.Role
	row := s.DB.QueryRow(`SELECT id, name FROM roles WHERE id = $1`, roleID)
	err := row.Scan(&role.ID, &role.Name)
	if err != nil {
		return nil, err
	}
	return &role, nil
}
