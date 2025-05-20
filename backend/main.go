package main

import (
    "log"                           // standard logging (used briefly for fallback)
    "net/http"                      // HTTP server and handler types
    "os"                            // for reading environment variables

    "github.com/adrian/gif-backend/handlers" // our HTTP handlers and middleware
    "github.com/gorilla/mux"                // request router
    "github.com/joho/godotenv"              // loads .env files into environment
    "github.com/sirupsen/logrus"            // structured, leveled logging
)

func main() {
    // 1) Load environment variables from a .env file, if present.
    //    If .env is missing, log a warning but continue (we expect vars to be set elsewhere).
    if err := godotenv.Load(); err != nil {
        logrus.Warn("No .env file found, relying on environment")
    }

    // 2) Ensure the GIPHY_API_KEY is set; this is critical for our service to work.
    //    If missing, we fatally exit (no sense running without an API key).
    if os.Getenv("GIPHY_API_KEY") == "" {
        logrus.Fatal("GIPHY_API_KEY is not set")
    }

    // 3) Configure Logrus to emit JSON-formatted logs for better parsing in production.
    logrus.SetFormatter(&logrus.JSONFormatter{})

    // 4) Create a new Gorilla Mux router to register routes and middleware.
    r := mux.NewRouter()

    // 5) Expose Prometheus metrics on the /metrics endpoint.
    //    handlers.ExposeMetricsHandler returns the promhttp.Handler.
    r.Handle("/metrics", handlers.ExposeMetricsHandler())

    // 6) Register our panic-recovery middleware first.
    //    It will catch any panics in downstream handlers, log them, and return a 500 response.
    r.Use(handlers.RecoveryMiddleware)

    // 7) CORS middleware: allow ONLY our React frontend origin to make requests.
    //    We use both the built-in Mux CORSMethodMiddleware and our custom handler.
    r.Use(mux.CORSMethodMiddleware(r))
    r.Use(corsMiddleware)

    // 8) Logging & metrics middleware: logs each request and updates Prometheus counters.
    r.Use(handlers.LoggingAndMetricsMiddleware)

    // 9) Health and readiness probes: simple JSON endpoints for uptime checks.
    r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
    r.HandleFunc("/ready", handlers.HealthCheck).Methods("GET")

    // 10) API routes are grouped under /api prefix.
    api := r.PathPrefix("/api").Subrouter()
    api.HandleFunc("/trending", handlers.GetTrending).Methods("GET")
    api.HandleFunc("/search", handlers.SearchGIFs).Methods("GET")

    // 11) Determine the port to listen on. Default to 5050 if PORT env var is missing.
    port := os.Getenv("PORT")
    if port == "" {
        port = "5050"
    }

    // 12) Log an info message indicating where the server is available.
    logrus.Infof("🚀 Backend running on http://localhost:%s", port)

    // 13) Start the HTTP server. If it fails, log.Fatal will exit the process.
    log.Fatal(http.ListenAndServe(":"+port, r))
}

// corsMiddleware sets CORS headers to allow cross-origin requests from our React app.
// It permits only GET and OPTIONS methods and the Content-Type header.
// For OPTIONS preflight requests, it returns immediately without calling the next handler.
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if r.Method == "OPTIONS" {
            // Preflight request: respond with headers only
            return
        }
        // For actual requests, proceed to the next handler
        next.ServeHTTP(w, r)
    })
}
