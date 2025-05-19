package main

import (
	"log"
	"net/http"
	"os"

	"github.com/adrian/gif-backend/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// load .env first
	if err := godotenv.Load(); err != nil {
		logrus.Warn("No .env file found, relying on environment")
	}

	// ensure API key is present
	if os.Getenv("GIPHY_API_KEY") == "" {
		logrus.Fatal("GIPHY_API_KEY is not set")
	}

	// use Logrus for structured JSON logs
	logrus.SetFormatter(&logrus.JSONFormatter{})

	r := mux.NewRouter()

	// metrics endpoint
	r.Handle("/metrics", handlers.ExposeMetricsHandler())

	// Panic-recovery middleware (must come before all others)
	r.Use(handlers.RecoveryMiddleware)

	// CORS Middleware
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(corsMiddleware)

	// Logging & metrics middleware
	r.Use(handlers.LoggingAndMetricsMiddleware)

	// health and readiness probes
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	r.HandleFunc("/ready", handlers.HealthCheck).Methods("GET")

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/trending", handlers.GetTrending).Methods("GET")
	api.HandleFunc("/search", handlers.SearchGIFs).Methods("GET")

	// start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}
	logrus.Infof("🚀 Backend running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
