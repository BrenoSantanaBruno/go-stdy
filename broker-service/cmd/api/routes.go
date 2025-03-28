package api

import (
	"net/http"

	"github.com/expl0iter/go-stdy/broker-service/internal/config"
	"github.com/expl0iter/go-stdy/broker-service/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Routes(cfg *config.Config) http.Handler {
	mux := chi.NewRouter()

	// Middleware
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Handlers
	brokerHandler := handlers.NewBrokerHandler(cfg)

	// Routes
	mux.Route("/api", func(r chi.Router) {
		r.Post("/broker", brokerHandler.Handle)
	})

	// Web routes
	mux.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/static/index.html")
	})

	return mux
}
