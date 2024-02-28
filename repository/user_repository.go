package repository

import (
	"database/sql"
	"kube/entity"
)

type IUserRepository interface {
	GetUserById(id string) (*entity.User, error)
	GetUserList() ([]entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserById(id string) (*entity.User, error) {
	user := entity.User{}
	err := ur.db.QueryRow("SELECT * FROM user WHERE id = ?", id).Scan(&user.Id, &user.Email, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) GetUserList() ([]entity.User, error) {
	rows, err := ur.db.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := []entity.User{}
	for rows.Next() {
		user := entity.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		res = append(res, user)
	}
	return res, nil
}

func (ur *userRepository) CreateUser(user *entity.User) (*entity.User, error) {
	_, err := ur.db.Exec(
		"INSERT INTO user (id, name, email) VALUES (?, ?, ?)",
	    user.Id,
		user.Name,
		user.Email,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) UpdateUser(user *entity.User) (*entity.User, error) {
	_, err := ur.db.Exec(
		"UPDATE user SET name = ?, email = ? WHERE id = ?",
		user.Name,
		user.Email,
	    user.Id,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
