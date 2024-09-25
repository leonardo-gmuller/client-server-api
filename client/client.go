package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Bid float64 `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("request to get exchange timed out.")
			panic(err)
		} else {
			log.Println(fmt.Errorf("send request to get exchange on https://economia.awesomeapi.com.br/json/last/USD-BRL failed: %w", err))
			panic(err)
		}
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(fmt.Errorf("read body failed: %w", err))
	}
	var data Response

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(fmt.Errorf("unmarshal json failed: %w", err))
	}
	saveExchangeTxt(data.Bid)
}

func saveExchangeTxt(bid float64) error {

	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Println(fmt.Errorf("create txt failed %w", err))
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar: %.2f", bid))
	if err != nil {
		log.Println(fmt.Errorf("write on txt failed %w", err))
		return err
	}
	return nil
}
