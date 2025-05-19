// File: backend/utils/giphy_test.go
package utils

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "os"
    "testing"
)

func TestFetchTrending(t *testing.T) {
    // Fake server with a single GIF
    fakeData := map[string]interface{}{
        "data": []map[string]interface{}{
            {"id": "abc123", "title": "Test GIF", "images": map[string]interface{}{
                "fixed_height": map[string]interface{}{"url": "https://example.com/1.gif"},
            }},
        },
        "pagination": map[string]interface{}{
            "total_count": 1, "count": 1, "offset": 0,
        },
    }
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/v1/gifs/trending" {
            t.Errorf("unexpected path: %s", r.URL.Path)
        }
        if key := r.URL.Query().Get("api_key"); key == "" {
            t.Error("expected api_key query param")
        }
        json.NewEncoder(w).Encode(fakeData)
    }))
    defer server.Close()

    // Override baseURL and set dummy key
    oldBase := baseURL
    baseURL = server.URL + "/v1/gifs"
    defer func() { baseURL = oldBase }()
    os.Setenv("GIPHY_API_KEY", "test-key")
    defer os.Unsetenv("GIPHY_API_KEY")

    // Call the typed function
    resp, err := FetchTrending(5, 1)
    if err != nil {
        t.Fatalf("FetchTrending error: %v", err)
    }

    // Now assert on the struct fields
    if len(resp.Data) != 1 {
        t.Fatalf("expected 1 GIF; got %d", len(resp.Data))
    }
    gif := resp.Data[0]
    if gif.ID != "abc123" {
        t.Errorf("expected ID 'abc123'; got %q", gif.ID)
    }
    if gif.Title != "Test GIF" {
        t.Errorf("expected Title 'Test GIF'; got %q", gif.Title)
    }
    if gif.Images.FixedHeight.URL != "https://example.com/1.gif" {
        t.Errorf("expected URL 'https://example.com/1.gif'; got %q", gif.Images.FixedHeight.URL)
    }

    // And check pagination
    if resp.Pagination.TotalCount != 1 {
        t.Errorf("expected total_count=1; got %d", resp.Pagination.TotalCount)
    }
}
