package handlers

import (
    "encoding/json" // encoding Go values to JSON
    "net/http" // HTTP request/response types
)

// HealthCheck responds with a simple JSON payload to indicate service health.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
    // 1. Set the Content-Type header so the client knows we're returning JSON.
    w.Header().Set("Content-Type", "application/json")
    // 2. Prepare the response body as a Go map. Here it's simple {"status":"ok"}.
    response := map[string]string{"status": "ok"}
    // 3. Encode the map to JSON and write it to the response writer.
    //    NewEncoder(w) wraps our ResponseWriter so json.Encoder writes directly to the HTTP response.
    //    Encode handles marshalling and adding a trailing newline.
    json.NewEncoder(w).Encode(response)
}