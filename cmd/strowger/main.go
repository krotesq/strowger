package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/krotesq/strowger/internal/account"
	"github.com/krotesq/strowger/internal/auth"
	"github.com/krotesq/strowger/internal/db"
	"github.com/krotesq/strowger/internal/mediamtx"
	"github.com/krotesq/strowger/internal/source"
	"github.com/krotesq/strowger/internal/target"
)

func main() {

	// connect to database
	ctx := context.Background()
	pool, err := db.Connect(ctx)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err.Error())
	}
	defer pool.Close()
	log.Println("connected to database")

	// create main router
	router := chi.NewRouter()

	// enable middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	// create api router
	routerApi := chi.NewRouter()

	// the account router has public and protected routes so we
	// mount the account router seperate and configure auth inside
	// we could move the /login, /register, /logout & /reset to the auth package sometime
	routerApi.Mount("/account", account.RoutesWithPool(pool))

	// protected routes
	routerApi.Group(func(r chi.Router) {
		r.Use(auth.Auth)
		r.Mount("/source", source.RoutesWithPool(pool))
		r.Mount("/target", target.RoutesWithPool(pool))
		r.Mount("/mediamtx", mediamtx.RoutesWithPool(pool))
	})

	// create web router
	routerWeb := chi.NewRouter()

	// add fs to web router
	webDir := filepath.Join(".", "web")
	fileServer := http.StripPrefix("/", http.FileServer(http.Dir(webDir)))
	routerWeb.Handle("/*", fileServer)

	// mount all routers to main router
	router.Mount("/api", routerApi)
	router.Mount("/", routerWeb)

	// run server
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	log.Printf("Server running at %s:%s", host, port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
