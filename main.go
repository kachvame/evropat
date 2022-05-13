package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.DefaultContextLogger = &log.Logger
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		NewLoggingMiddleware(log.Logger.With().Str("component", "web").Logger()),
		middleware.Recoverer,
		middleware.Timeout(60*time.Second),
	)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/cities", func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, r, citiesResponse{
				Cities: cities,
			})
		})
	})

	addr := ":3333"
	log.Printf("slusham i kachvam na port %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		return err
	}
	return nil
}
