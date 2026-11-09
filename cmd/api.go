package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/oskarscot/shortly/internal/link"
)

func (a *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	linkRepo := link.NewMemoryRepository()
	linkSvc := link.NewService(linkRepo)
	linkHandler := link.NewHandler(*linkSvc)

	linkHandler.Setup(r)

	return r
}

func (a *application) run(r http.Handler) error {
	srv := &http.Server{
		Handler:      r,
		Addr:         a.config.address,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	return srv.ListenAndServe()
}

type application struct {
	config   configuration
	dbConfig dbConfig
}

type configuration struct {
	address string
}

type dbConfig struct {
	dsn string
}
