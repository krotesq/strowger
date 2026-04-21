package account

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RoutesWithPool(pool *pgxpool.Pool) chi.Router {
	
	router := chi.NewRouter()
	
	repository := newRepository(pool)
	service := newService(repository)
	handler := newHandler(service)
	
	router.Get("/account", handler.findByUsername)
	router.Post("/login", handler.login)
	
	return router
}