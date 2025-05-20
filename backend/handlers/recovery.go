package handlers

import (
    "encoding/json"       // JSON encoding for error responses
    "net/http"            // HTTP types for handlers and status codes

    "github.com/sirupsen/logrus" // Structured logging
)

// RecoveryMiddleware wraps HTTP handlers to catch any panics (runtime errors),
// log them, and return a clean 500 Internal Server Error instead of crashing.
func RecoveryMiddleware(next http.Handler) http.Handler {
    // Return a new handler that adds recovery around the original one.
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Defer a function that will run if a panic occurs below.
        defer func() {
            if rec := recover(); rec != nil {
                // 1) Log the panic value for diagnosis.
                logrus.WithField("panic", rec).Error("panic recovered in handler")

                // 2) Prepare a JSON error response for the client.
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusInternalServerError)
                // 3) Encode a simple {"error":"internal server error"} payload.
                json.NewEncoder(w).Encode(map[string]string{
                    "error": "internal server error",
                })
            }
        }()
        // Call the next handler in the chain. If it panics, our defer will catch it.
        next.ServeHTTP(w, r)
    })
}
