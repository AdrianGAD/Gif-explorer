package handlers

import (
    "encoding/json"                   // for encoding Go values to JSON
    "net/http"                        // for HTTP request and response types
    "os"                              // to access environment variables
    "strconv"                         // for converting strings to integers

    "github.com/adrian/gif-backend/utils" // our internal package for Giphy API calls
)

// SearchGIFs handles GET /api/search requests. It reads query parameters,
// calls the Giphy search utility, and returns a JSON payload of GIFs.
func SearchGIFs(w http.ResponseWriter, r *http.Request) {
    // 1) Extract query parameters from the URL
    q := r.URL.Query().Get("q")         // search term (required)
    limit := r.URL.Query().Get("limit") // number of items per page
    page := r.URL.Query().Get("page")   // pagination page number
    rating := r.URL.Query().Get("rating") // content rating filter (optional)

    // 2) Validate required parameters
    if q == "" {
        // If 'q' is missing, return 400 Bad Request with an error message
        http.Error(w, "Query param 'q' is required", http.StatusBadRequest)
        return
    }

    // 3) Provide default values if optional params are omitted
    if limit == "" {
        limit = "12"
    }
    if page == "" {
        page = "1"
    }

    // 4) Convert limit and page to integers, ignoring errors (defaults apply)
    limitInt, _ := strconv.Atoi(limit)
    pageInt, _ := strconv.Atoi(page)

    // 5) Ensure the GIPHY_API_KEY is set in the environment
    if os.Getenv("GIPHY_API_KEY") == "" {
        http.Error(w, "Server misconfigured: missing GIPHY_API_KEY", http.StatusInternalServerError)
        return
    }

    // 6) Call the typed SearchGIFs utility, which returns a GiphyResponse struct
    result, err := utils.SearchGIFs(q, rating, limitInt, pageInt)
    if err != nil {
        // If the Giphy API call fails, return 500 Internal Server Error
        http.Error(w, "Failed to fetch search results", http.StatusInternalServerError)
        return
    }

    // 7) Write the successful JSON response
    w.Header().Set("Content-Type", "application/json") // tell the client it’s JSON
    // Encode the GiphyResponse directly to the HTTP response body
    json.NewEncoder(w).Encode(result)
}
