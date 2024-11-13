package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/cachedb"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/cachemodel"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/datamodel"
)

type CoinEntityRepository struct {
	db *mysql.ApplicationDB
}

func NewCoinEntityRepository(db *mysql.ApplicationDB) *CoinEntityRepository {
	return &CoinEntityRepository{db: db}
}

func (r *CoinEntityRepository) BulkInsert(ctx context.Context, tx *sqlx.Tx, contents []cachedb.CacheContent) error {
	dms, err := r.toDataModels(contents)
	if err != nil {
		return err
	}

	err = dms.InsertAll(ctx, tx, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (r *CoinEntityRepository) BulkUpdate(ctx context.Context, tx *sqlx.Tx, contents []cachedb.CacheContent) error {
	dms, err := r.toDataModels(contents)
	if err != nil {
		return err
	}
	query := "UPDATE user_coin SET" + r.createUpdateCaseQuery(dms) + "WHERE" + r.createUpdateWhereQuery(dms)
	if _, err = tx.ExecContext(ctx, query); err != nil {
		return err
	}
	return nil
}

func (r *CoinEntityRepository) BulkDelete(ctx context.Context, tx *sqlx.Tx, content []cachedb.CacheContent) error {

	return nil
}

func (r *CoinEntityRepository) toDataModels(contents []cachedb.CacheContent) (datamodel.UserCoinSlice, error) {
	var err error
	dms := lo.Map(contents, func(content cachedb.CacheContent, _ int) *datamodel.UserCoin {
		m, ok := content.(*cachemodel.UserCoinCacheModel)
		if !ok {
			err = errors.New("failed to convert to datamodel")
		}

		return &datamodel.UserCoin{
			UserID: m.UserID,
			Num:    m.Num,
		}
	})
	return dms, err
}

func (r *CoinEntityRepository) createUpdateCaseQuery(dms []*datamodel.UserCoin) string {
	query := "num = case "
	for _, dm := range dms {
		condition := fmt.Sprintf("WHEN user_id=%d THEN %d ", dm.UserID, dm.Num)
		query += condition
	}
	query += "Else 0 End"
	return query
}

func (r *CoinEntityRepository) createUpdateWhereQuery(dms []*datamodel.UserCoin) string {
	query := ""
	for index, dm := range dms {
		condition := fmt.Sprintf("(user_id=%d)", dm.UserID)
		if index != len(dms)-1 {
			condition += " OR "
		}
		query += condition
	}
	return query
}
