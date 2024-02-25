package repository

import "kube/entity"

type IUserRepository interface {
	GetUser(id string) (*entity.User, error)
}

type userRepository struct {}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (ur *userRepository) GetUser(id string) (*entity.User, error) {
	user := entity.User{Id: id, Email: "karasawa@gmail.com", Name: "karasawa takunon"}
	return &user, nil
}
