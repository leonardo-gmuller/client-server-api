package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/leonardo-gmulller/client-server-api/server/database/repositories"
	"github.com/leonardo-gmulller/client-server-api/server/dto"
)

type Handler struct {
	ExchangeRepository repositories.ExchangeRepository
}

type ResponseClient struct {
	Bid float64 `json:"bid"`
}

func Init(db *sql.DB) {
	handler := Handler{
		ExchangeRepository: *repositories.NewExchangeRepository(db),
	}
	http.HandleFunc("/cotacao", handler.requestExchangeHandler)
	http.ListenAndServe(":8080", nil)
}

func (h *Handler) requestExchangeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	exchange, err := requestExchange(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.ExchangeRepository.CreateExchange(ctx, exchange)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bid, err := strconv.ParseFloat(exchange.Usdbrl.Bid, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := ResponseClient{
		Bid: bid,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func requestExchange(ctx context.Context) (*dto.RequestExchange, error) {
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Println(fmt.Errorf("prepare request with context failed: %w", err))
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("request to get exchange on https://economia.awesomeapi.com.br/json/last/USD-BRL timed out.")
			return nil, err
		} else {
			log.Println(fmt.Errorf("send request to get exchange on https://economia.awesomeapi.com.br/json/last/USD-BRL failed: %w", err))
			return nil, err
		}
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(fmt.Errorf("read body failed: %w", err))
		return nil, err
	}
	var data dto.RequestExchange

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(fmt.Errorf("unmarshal json failed: %w", err))
		return nil, err
	}

	return &data, nil
}
