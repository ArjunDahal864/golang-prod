package service

import (
	"prod/dao"
	"prod/db"
	repo "prod/repository"
)

func NewUserService(repo *repo.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}

}

type UserService struct {
	repo *repo.UserRepo
}

func (u *UserService) CreateUser(user *dao.UserRequestDao) (*dao.UserResponseDao, error) {
	usr := db.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}

	dbUser,err := u.repo.Create(&usr)
	if err != nil {
		return nil, err
	}
	responseUser := dao.UserResponseDao{
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Email:     dbUser.Email,
	}
	return &responseUser, nil
}
