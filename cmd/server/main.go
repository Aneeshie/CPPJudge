package main

import (
	"context"
	"log"

	"github.com/Aneeshie/cpp-judge/internal/config"
	"github.com/Aneeshie/cpp-judge/internal/database"
)

func main(){
	cfg := config.Load()
	pool, err := database.NewPool(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

}
