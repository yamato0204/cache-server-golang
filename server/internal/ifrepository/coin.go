package ifrepository

import (
	"context"

	"github.com/yamato0204/cache-server-golang/internal/entity"
)

type IFCoinEntityRepository interface {
	Create(ctx context.Context, userId int64) (*entity.CoinEntity, error)
	Save(ctx context.Context, e *entity.CoinEntity) error
}
