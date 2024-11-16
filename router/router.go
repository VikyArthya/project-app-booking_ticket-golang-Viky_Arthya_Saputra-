package router

import (
	"ticket-booking/handler"
	"ticket-booking/middleware"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// Middleware untuk logging request
	r.Use(middleware.RequestLogger)

	// Routes
	r.Post("/booking", handler.CreateBookingHandler)
	r.Get("/booking/{id}", handler.GetBookingByIDHandler)
	r.Get("/booking", handler.GetAllBookingsHandler)
	r.Put("/booking/{id}", handler.UpdateBookingHandler)
	r.Delete("/booking/{id}", handler.DeleteBookingHandler)

	return r
}
