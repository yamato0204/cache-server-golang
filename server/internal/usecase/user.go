package usecase

import (
	"context"

	"github.com/yamato0204/cache-server-golang/internal/service"
)

type UserUsercase struct {
	userRegisterService *service.UserRegisterService
}

func NewUserUsecase(userRegisterService *service.UserRegisterService) *UserUsercase {
	return &UserUsercase{
		userRegisterService: userRegisterService,
	}
}

func (u *UserUsercase) Register(ctx context.Context) (int64, error) {
	userId, err := u.userRegisterService.Register(ctx)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
