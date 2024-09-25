package repositories

import "database/sql"

type ExchangeRepository struct {
	Db *sql.DB
}

func NewExchangeRepository(db *sql.DB) *ExchangeRepository {
	return &ExchangeRepository{
		Db: db,
	}
}
