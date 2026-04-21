package main

import (
	"context"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/krotesq/strowger/internal/account"
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

	// attach sub routers
	routerApi.Mount("/account", account.RoutesWithPool(pool))
	routerApi.Mount("/source", source.RoutesWithPool(pool))
	routerApi.Mount("/target", target.RoutesWithPool(pool))
	routerApi.Mount("/mediamtx", mediamtx.RoutesWithPool(pool))
	
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
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
