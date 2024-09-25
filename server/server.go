package main

import (
	"fmt"

	"github.com/leonardo-gmulller/client-server-api/server/database"
	"github.com/leonardo-gmulller/client-server-api/server/handler"
)

func main() {
	DB, err := database.Init()
	if err != nil {
		fmt.Println(err)
	}
	defer DB.Close()

	handler.Init(DB)
}
