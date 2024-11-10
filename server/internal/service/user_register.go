package service

import (
	"context"

	"github.com/yamato0204/cache-server-golang/internal/ifrepository"
)

type UserRegisterService struct {
	coinEntityRepository ifrepository.IFCoinEntityRepository
}

func NewUserRegisterService(coinEntityRepository ifrepository.IFCoinEntityRepository) *UserRegisterService {
	return &UserRegisterService{coinEntityRepository: coinEntityRepository}
}

func (s *UserRegisterService) Register(ctx context.Context) (int64, error) {
	//userId := s.uuidGenerator.GetUUID()

	if err := s.addCoin(ctx, 222); err != nil {
		return 0, err
	}

	return 0, nil
}

func (s *UserRegisterService) addCoin(ctx context.Context, userId int64) error {
	e, err := s.coinEntityRepository.Create(ctx, userId)
	if err != nil {
		return err
	}
	e.Add(1000)

	return nil
}
