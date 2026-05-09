package main

import "log"

func main() {

	config := config{
		addr: ":8080",
	}

	app := &application{config: config}

	log.Fatal(app.run)
}
