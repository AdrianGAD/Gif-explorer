package utils

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "os"
    "testing"
)

func TestFetchTrending(t *testing.T) {
    // 1) Spin up a fake Giphy server that expects an api_key
    fakeData := map[string]interface{}{
        "data": []map[string]interface{}{
            {"id": "abc123", "type": "gif"},
        },
    }
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // verify path
        if r.URL.Path != "/v1/gifs/trending" {
            t.Errorf("unexpected path: %s", r.URL.Path)
        }
        // verify non-empty api_key
        if key := r.URL.Query().Get("api_key"); key == "" {
            t.Error("expected api_key query param")
        }
        // return fake JSON
        json.NewEncoder(w).Encode(fakeData)
    }))
    defer server.Close()

    // 2) Override baseURL so our code calls the fake server
    oldBase := baseURL
    baseURL = server.URL + "/v1/gifs"
    defer func() { baseURL = oldBase }()

    // 3) **Set a dummy GIPHY_API_KEY** so the test sees a non-empty api_key
    os.Setenv("GIPHY_API_KEY", "test-key")
    defer os.Unsetenv("GIPHY_API_KEY")

    // 4) Call the function under test
    result, err := FetchTrending(5, 1)
    if err != nil {
        t.Fatalf("FetchTrending returned error: %v", err)
    }

    // 5) Assert we got back our fake data
    gotData, ok := result["data"].([]interface{})
    if !ok || len(gotData) != 1 {
        t.Fatalf("expected data length 1; got %#v", result["data"])
    }
    first := gotData[0].(map[string]interface{})
    if first["id"] != "abc123" {
        t.Errorf("unexpected id: %v", first["id"])
    }
}
