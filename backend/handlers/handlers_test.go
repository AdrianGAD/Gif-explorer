package handlers

import (
    "net/http"           // for HTTP status codes and method constants
    "net/http/httptest"  // to create fake Request and ResponseRecorder
    "strings"            // for simple substring checks in response bodies
    "testing"            // the Go testing framework
)

// TestHealthCheck verifies that the HealthCheck handler returns a 200 status
// and the exact JSON payload {"status":"ok"}.
func TestHealthCheck(t *testing.T) {
    // Create a new GET request targeting the /health endpoint.
    // The body is nil because HealthCheck doesn't read from it.
    req := httptest.NewRequest(http.MethodGet, "/health", nil)

    // httptest.NewRecorder gives us a ResponseRecorder to capture the handler's output.
    w := httptest.NewRecorder()

    // Call the HealthCheck handler directly, passing in our fake request & recorder.
    HealthCheck(w, req)

    // Obtain the result from the recorder.
    resp := w.Result()
    defer resp.Body.Close() // ensure we close the body when done

    // 1) Check the HTTP status code.
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("expected status 200 OK; got %d", resp.StatusCode)
    }

    // 2) Read the response body as a string.
    body := w.Body.String()

    // The handler is expected to write exactly {"status":"ok"} (no trailing newline).
    expected := `{"status":"ok"}`
    if strings.TrimSpace(body) != expected {
        t.Errorf("unexpected body:\n got %q\nwant %q", body, expected)
    }
}

// TestSearchGIFs_MissingQuery verifies that SearchGIFs returns a 400 Bad Request
// when the required 'q' parameter is absent.
func TestSearchGIFs_MissingQuery(t *testing.T) {
    // Create a GET request to /api/search without any query string.
    req := httptest.NewRequest(http.MethodGet, "/api/search", nil)
    w := httptest.NewRecorder()

    // Invoke the SearchGIFs handler.
    SearchGIFs(w, req)

    resp := w.Result()
    defer resp.Body.Close()

    // Expect a 400 Bad Request status code when 'q' is missing.
    if resp.StatusCode != http.StatusBadRequest {
        t.Fatalf("expected status 400 Bad Request; got %d", resp.StatusCode)
    }

    // Check that the error message mentions the missing 'q' parameter.
    body := w.Body.String()
    if !strings.Contains(body, "Query param 'q' is required") {
        t.Errorf("expected error about missing q; got %q", body)
    }
}

// You can add more tests below. For example:
// - TestSearchGIFs_Success: stub out http.Get via a custom http.Client in tests
//   and verify that valid ?q=cat returns forwarded JSON from your stub.
// - TestGetTrending: similarly stub utils.FetchTrending or http.Get to return
//   canned data, then confirm GetTrending writes that data unmodified.
// - Table-driven tests for different limit/page values, invalid numbers, etc.
