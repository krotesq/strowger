package account

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RoutesWithPool(pool *pgxpool.Pool) chi.Router {
	router := chi.NewRouter()
	return router
}