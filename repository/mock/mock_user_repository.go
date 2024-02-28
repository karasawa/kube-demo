package mock

import (
	"kube/entity"
	"kube/repository"
)

type mockUserRepository struct{}

func NewMockUserRepository() repository.IUserRepository {
	return &mockUserRepository{}
}

func (mur *mockUserRepository) GetUserById(id string) (*entity.User, error) {
	user := entity.User{Id: id, Email: "test@gmail.com", Name: "test AAAA"}
	return &user, nil
}

func (mur *mockUserRepository) GetUserList() ([]entity.User, error) {
	res := []entity.User{}
	user1 := entity.User{Id: "100", Email: "test1@gmail.com", Name: "test AAAA"}
	user2 := entity.User{Id: "200", Email: "test2@gmail.com", Name: "test BBBB"}
	res = append(res, user1)
	res = append(res, user2)
	return res, nil
}

func (mur *mockUserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	return user, nil
}

func (mur *mockUserRepository) UpdateUser(user *entity.User) (*entity.User, error) {
	return user, nil
}
