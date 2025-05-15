package main

import (
	"log"
	"net/http"
	"os"

	"github.com/adrian/gif-backend/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	r := mux.NewRouter()

	// CORS Middleware
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(corsMiddleware)

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/trending", handlers.GetTrending).Methods("GET")
	api.HandleFunc("/search", handlers.SearchGIFs).Methods("GET")
	r.HandleFunc("/api/search", handlers.SearchGIFs).Methods("GET")
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}

	log.Printf("🚀 Backend running on http://localhost:%s\n", port)
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
