package integrationrepository

import (
	"context"

	"github.com/yamato0204/cache-server-golang/internal/entity"
	"github.com/yamato0204/cache-server-golang/internal/ifrepository"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/cacherepository"
)

type CoinIntegrationRepository struct {
	coinRepository ifrepository.IFCoinEntityRepository
}

func NewCoinIntegrationRepository(cacheRepository *cacherepository.CoinEntityCacheRepository) *CoinIntegrationRepository {
	return &CoinIntegrationRepository{coinRepository: cacheRepository}

}

func (r *CoinIntegrationRepository) Create(ctx context.Context, userId int64) (*entity.CoinEntity, error) {
	return r.coinRepository.Create(ctx, userId)
}

func (r *CoinIntegrationRepository) Save(ctx context.Context, e *entity.CoinEntity) error {
	return r.coinRepository.Save(ctx, e)
}
