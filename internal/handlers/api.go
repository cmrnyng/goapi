package handlers

import (
	"github.com/cmrnyng/goapi/internal/handlers/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) { // Uppercase function name tells compiler that function can be imported
													 // in other packages. If it's lowercase, it will be a private function.
	// Global middleware - global meaning it is applied to all endpoints
	r.Use(chimiddle.StripSlashes) // StripSlashes removes trailing slashes

	r.Route("/account", func(router chi.Router) {
		// Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}