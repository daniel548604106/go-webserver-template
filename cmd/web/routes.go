package main

import (
	"net/http"

	"github.com/daniel548604106/go-bookings/internal/config"
	"github.com/daniel548604106/go-bookings/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// CSRF protection
	// Server generates a unique token per session/request.
	// Token is stored server-side and sent to the client (cookie + form field).
	// Forms include the token as a hidden field.
	// On POST/PUT/DELETE, the server verifies the token matches.
	// If it doesn't match, the request is rejected.
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Post("/search-availability-json", handlers.Repo.AvailibilityJSON)
	mux.Get("/search-availibility", handlers.Repo.Availability)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)
	mux.Post("/search-availibility", handlers.Repo.PostAvailability)

	// Serve static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
