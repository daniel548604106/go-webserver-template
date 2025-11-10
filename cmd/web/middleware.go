package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf is the csrf protection middleware
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	// Generates a unique token per session
	// Stores it in a cookie (HttpOnly, Secure, SameSite)
	// Validates it on state-changing requests
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/", // Available site-wide
		Secure:   app.InProduction,  // HTTPS only in production
		SameSite: http.SameSiteLaxMode,   // Prevents cross-site requests
	})
	return csrfHandler
}

// SessionLoad loads and saves session data for current request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
