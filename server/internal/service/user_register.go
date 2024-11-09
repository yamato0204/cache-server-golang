package service

import "context"

type UserRegisterService struct {
	uuidGenerator        ifrepository.IFUUIDGenerator
	coinEntityRepository ifrepository.IFCoinEntityRepository
}

func NewUserRegisterService(
	uuidGenerator ifrepository.IFUUIDGenerator,
	coinEntityRepository ifrepository.IFCoinEntityRepository,
) *UserRegisterService {
	return &UserRegisterService{
		uuidGenerator:        uuidGenerator,
		coinEntityRepository: coinEntityRepository,
	}
}

func (s *UserRegisterService) Register(ctx context.Context) (int64, error) {
	userId := s.uuidGenerator.GetUUID()

	if err := s.addCoin(ctx, userId); err != nil {
		return 0, err
	}

	return userId, nil
}

func (s *UserRegisterService) addCoin(ctx context.Context, userId int64) error {
	e, err := s.coinEntityRepository.Create(ctx, userId)
	if err != nil {
		return err
	}
	e.Add(1000)
	return s.coinEntityRepository.Save(ctx, e)
}
