package main

import (
	"context"
	"log"

	"github.com/Aneeshie/cpp-judge/internal/config"
	"github.com/Aneeshie/cpp-judge/internal/database"
	"github.com/Aneeshie/cpp-judge/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env found")
	}

	cfg := config.Load()
	pool, err := database.NewPool(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	svr := server.NewServer(pool)

	svr.Run(cfg.Port)

}
