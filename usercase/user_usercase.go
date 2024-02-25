package usercase

import (
	"kube/entity"
	"kube/repository"
)

type IUserUsecase interface {
	GetUser(id string) (*entity.User, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) GetUser(id string) (*entity.User, error) {
	res, err := uu.ur.GetUser(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
