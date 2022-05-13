package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"regexp"
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
		render.SetContentType(render.ContentTypeJSON),
	)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/cities", func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, r, citiesResponse{
				Cities: cities,
			})
		})
		r.Get("/offices", func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, r, officesResponse{
				Offices: offices,
			})
		})
	})

	r.Route("/v6", func(r chi.Router) {
		r.Post("/waybills", func(w http.ResponseWriter, r *http.Request) {
			if rand.Float64() < 0.23 {
				panic("opa") // todo error handling
			}

			dto := shippingRequest{}
			if err := render.DecodeJSON(r.Body, &dto); err != nil {
				render.JSON(w, r, errorResponse{
					Error: err,
				})
				return
			}

			shippingID := sha256.Sum256([]byte(fmt.Sprintf("%d-%d", dto.CityID, dto.OfficeID)))
			render.JSON(w, r, shippingResponse{
				ID:    hex.EncodeToString(shippingID[:]),
				Extra: regexp.MustCompile(`\w+`).ReplaceAllString(dto.Extra, "[censored]"),
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
