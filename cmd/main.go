package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	log.SetFlags(log.LstdFlags)
	if err := start(); err != nil {
		log.Fatal(err)
	}
}

func start() error {
	dbConfig := dbConfig{
		dsn: os.Getenv("GOOSE_DBSTRING"),
	}

	config := configuration{
		address: os.Getenv("APP_URL"),
	}

	app := application{
		config:   config,
		dbConfig: dbConfig,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgxpool.New(ctx, dbConfig.dsn)
	if err != nil {
		return err
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		return err
	}

	log.Printf("server starting at %s\n", app.config.address)

	return app.run(app.mount(pool))
}
