package main

import "log"

func main() {
	log.SetFlags(log.LstdFlags)

	config := configuration{
		address: ":8080",
	}

	app := application{
		config: config,
	}

	log.Printf("server starting at %s\n", app.config.address)

	if err := app.run(app.mount()); err != nil {
		log.Fatal(err)
	}
}
