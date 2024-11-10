package entity

type CoinEntity struct {
	useId int64
	num   int64
}

func NewCoinEntity(userId int64) *CoinEntity {
	return &CoinEntity{
		useId: userId,
		num:   0,
	}
}

func (e *CoinEntity) Add(num int64) {
	e.num += num
}

func (e *CoinEntity) UserId() int64 {
	return e.useId
}

func (e *CoinEntity) Num() int64 {
	return e.num
}
