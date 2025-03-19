package routes

import (
	"expense-tracker-api/internal/handlers"
	m "expense-tracker-api/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/signup", handlers.Signup)
	r.Post("/login", handlers.Login)

	r.Route("/expenses", func(r chi.Router) {
		r.Use(m.AuthMiddleware)
		r.Post("/", handlers.AddExpense)
		r.Get("/", handlers.ListExpenses)
	})

	return r
}
