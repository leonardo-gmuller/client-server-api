package repositories

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/leonardo-gmulller/client-server-api/server/dto"
	"github.com/leonardo-gmulller/client-server-api/server/entity"
)

func (repository *ExchangeRepository) CreateExchange(ctx context.Context, exchange *dto.RequestExchange) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	stmt, err := repository.Db.Prepare("insert into exchange_rate (price, created_at) values(?, ?)")

	if err != nil {
		log.Println(fmt.Errorf("prepare statement failed: %w", err))
		return err
	}
	bid, err := strconv.ParseFloat(exchange.Usdbrl.Bid, 64)
	if err != nil {
		log.Println(fmt.Errorf("prepare request with context failed: %w", err))
		return err
	}
	exchangeRate := entity.Exchange{
		Price:      bid,
		Created_at: time.Now().String(),
	}

	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, exchangeRate.Price, exchangeRate.Created_at)

	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Insert operation timed out.")
			return err
		} else {
			log.Println(fmt.Errorf("execute statement failed: %w", err))
			return err
		}
	}
	return nil

}
