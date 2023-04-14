package usecase

import (
	"github.com/febriansr/jwt-auth/model"
	"github.com/febriansr/jwt-auth/repository"
)

type UserUsecase interface {
	Login(inputedUser *model.User) any
}

type userUsecase struct {
	userRepo repository.UserRepo
}

func (u *userUsecase) Login(inputedUser *model.User) any {
	return u.userRepo.Login(inputedUser)
}

func NewUserUsecase(userRepo repository.UserRepo) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
