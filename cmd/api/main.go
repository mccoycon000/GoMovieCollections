package main

import (
	"log"
	"movie_collections/internal/env"
	"movie_collections/internal/store"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
