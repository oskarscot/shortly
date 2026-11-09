package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags)

	dbConfig := dbConfig{
		dsn: "",
	}

	config := configuration{
		address: os.Getenv("APP_URL"),
	}

	app := application{
		config:   config,
		dbConfig: dbConfig,
	}

	log.Printf("server starting at %s\n", app.config.address)

	if err := app.run(app.mount()); err != nil {
		log.Fatal(err)
	}
}
