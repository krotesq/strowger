package account

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/krotesq/strowger/internal/auth"
)

func RoutesWithPool(pool *pgxpool.Pool) chi.Router {

	router := chi.NewRouter()

	repository := newRepository(pool)
	service := newService(repository)
	handler := newHandler(service)

	// public
	router.Post("/login", handler.login)
	router.Get("/logout", handler.logout)

	// protected
	router.Group(func(r chi.Router) {
		r.Use(auth.Auth)
		r.Post("/", handler.create)
		r.Get("/{id}", handler.findByID)
		r.Patch("/{id}/deactivate", handler.deactivateByID)
		r.Patch("/{id}/activate", handler.activateByID)
		r.Delete("/{id}", handler.deleteByID)
		r.Get("/me", handler.me)
	})

	return router
}
