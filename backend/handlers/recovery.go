package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/sirupsen/logrus"
)

// recoveryMiddleware catches panics in handlers, logs them, and returns a 500 JSON error.
func RecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if rec := recover(); rec != nil {
                // Log the panic with its value
                logrus.WithField("panic", rec).Error("panic recovered in handler")

                // Return a generic 500 response
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusInternalServerError)
                json.NewEncoder(w).Encode(map[string]string{
                    "error": "internal server error",
                })
            }
        }()
        next.ServeHTTP(w, r)
    })
}
