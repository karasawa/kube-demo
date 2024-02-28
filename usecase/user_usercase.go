package usecase

import (
	"kube/entity"
	"kube/repository"
	"kube/util"
)

type IUserUsecase interface {
	GetUserById(id string) (*entity.User, error)
	GetUserList() ([]entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) GetUserById(id string) (*entity.User, error) {
	res, err := uu.ur.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uu *userUsecase) GetUserList() ([]entity.User, error) {
	res, err := uu.ur.GetUserList()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uu *userUsecase) CreateUser(user *entity.User) (*entity.User, error) {
	user.Id = util.CreateUuid()
	res, err := uu.ur.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uu *userUsecase) UpdateUser(user *entity.User) (*entity.User, error) {
	res, err := uu.ur.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}