package handlers

import (
    "encoding/json"                   // JSON encoding for responses
    "net/http"                        // HTTP request/response types
    "os"                              // for reading environment variables
    "strconv"                         // for converting strings to integers

    "github.com/adrian/gif-backend/utils" // our internal utils for fetching GIF data
)

// GetTrending handles GET requests to /api/trending.
// It reads pagination parameters, calls the utils.FetchTrending function,
// and writes a JSON response containing the trending GIFs.
func GetTrending(w http.ResponseWriter, r *http.Request) {
    // 1) Parse optional query parameters "limit" and "page"
    //    Default to 12 items per page and page 1 if not provided.
    limit := r.URL.Query().Get("limit")
    page  := r.URL.Query().Get("page")
    if limit == "" {
        limit = "12"
    }
    if page == "" {
        page = "1"
    }

    // 2) Convert limit and page from strings to integers.
    //    We ignore errors here, falling back to defaults above.
    limitInt, _ := strconv.Atoi(limit)
    pageInt,  _ := strconv.Atoi(page)
    
    // 3) Check that the GIPHY_API_KEY environment variable is set.
    //    If it's missing, we cannot call Giphy, so return a 500 error.
    if os.Getenv("GIPHY_API_KEY") == "" {
        http.Error(w, "Server misconfigured: missing GIPHY_API_KEY", http.StatusInternalServerError)
        return
    }

    // 4) Call the typed FetchTrending utility with our pagination values.
    //    This returns a GiphyResponse struct or an error.
    respData, err := utils.FetchTrending(limitInt, pageInt)
    if err != nil {
        // If fetching from Giphy fails, return a 500 error to the client.
        http.Error(w, "Failed to fetch trending GIFs", http.StatusInternalServerError)
        return
    }

    // 5) On success, set the Content-Type header to application/json.
    w.Header().Set("Content-Type", "application/json")
    // 6) Encode the GiphyResponse struct directly to the HTTP response body.
    //    json.NewEncoder(w) writes the JSON and a trailing newline.
    json.NewEncoder(w).Encode(respData)
}
