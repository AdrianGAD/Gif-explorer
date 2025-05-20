package utils

import (
    "encoding/json"       // to encode our fake data as JSON
    "net/http"            // for HTTP status and request types
    "net/http/httptest"   // to create a fake HTTP server
    "os"                  // to set/unset environment variables
    "testing"             // Go’s testing framework
)

// TestFetchTrending verifies that FetchTrending correctly calls the Giphy API,
// passes along the API key, and parses the JSON into our typed GiphyResponse.
func TestFetchTrending(t *testing.T) {
    // 1) Prepare fake Giphy JSON payload with one GIF and pagination info.
    fakeData := map[string]interface{}{
        "data": []map[string]interface{}{ // top-level "data" array
            {
                "id":    "abc123",
                "title": "Test GIF",
                "images": map[string]interface{}{ // nested "images" object
                    "fixed_height": map[string]interface{}{ // only test the fixed_height variant
                        "url": "https://example.com/1.gif",
                    },
                },
            },
        },
        "pagination": map[string]interface{}{ // top-level "pagination" object
            "total_count": 1,
            "count":       1,
            "offset":      0,
        },
    }

    // 2) Start an httptest.Server to simulate Giphy’s API.
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 2a) Verify the request path matches our expected endpoint.
        if r.URL.Path != "/v1/gifs/trending" {
            t.Errorf("unexpected path: %s", r.URL.Path)
        }
        // 2b) Ensure the API key query parameter is non-empty.
        if key := r.URL.Query().Get("api_key"); key == "" {
            t.Error("expected api_key query param")
        }
        // 2c) Write our fake JSON payload to the response.
        json.NewEncoder(w).Encode(fakeData)
    }))
    defer server.Close() // shut down server when test completes

    // 3) Override the package-level baseURL so FetchTrending points at our fake server.
    oldBase := baseURL
    baseURL = server.URL + "/v1/gifs"
    defer func() { baseURL = oldBase }() // restore original baseURL after test

    // 4) Set a dummy GIPHY_API_KEY so FetchTrending includes it in requests.
    os.Setenv("GIPHY_API_KEY", "test-key")
    defer os.Unsetenv("GIPHY_API_KEY") // clean up afterwards

    // 5) Call the function under test
    resp, err := FetchTrending(5, 1)
    if err != nil {
        t.Fatalf("FetchTrending error: %v", err)
    }

    // 6) Assert that one GIF was returned
    if len(resp.Data) != 1 {
        t.Fatalf("expected 1 GIF; got %d", len(resp.Data))
    }
    gif := resp.Data[0]

    // 7) Validate that the fields were parsed correctly
    if gif.ID != "abc123" {
        t.Errorf("expected ID 'abc123'; got %q", gif.ID)
    }
    if gif.Title != "Test GIF" {
        t.Errorf("expected Title 'Test GIF'; got %q", gif.Title)
    }
    if gif.Images.FixedHeight.URL != "https://example.com/1.gif" {
        t.Errorf(
            "expected URL 'https://example.com/1.gif'; got %q",
            gif.Images.FixedHeight.URL,
        )
    }

    // 8) Verify pagination metadata
    if resp.Pagination.TotalCount != 1 {
        t.Errorf("expected total_count=1; got %d", resp.Pagination.TotalCount)
    }
}
