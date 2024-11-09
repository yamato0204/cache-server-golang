package service

import (
	"context"
	"fmt"
)

type UserRegisterService struct {
	userId int64
}

func NewUserRegisterService(
	userId int64,
) *UserRegisterService {
	return &UserRegisterService{
		userId: userId,
	}
}

func (s *UserRegisterService) Register(ctx context.Context) (int64, error) {
	// userId := s.uuidGenerator.GetUUID()

	// if err := s.addCoin(ctx, userId); err != nil {
	// 	return 0, err
	// }
	fmt.Println(s.userId)

	return 0, nil
}

func (s *UserRegisterService) addCoin(ctx context.Context, userId int64) error {
	// e, err := s.coinEntityRepository.Create(ctx, userId)
	// if err != nil {
	// 	return err
	// }
	// e.Add(1000)

	return nil
}
