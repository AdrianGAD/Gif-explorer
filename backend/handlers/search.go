package handlers

import (
    "encoding/json"
    "net/http"
    "os"
    "strconv"

    "github.com/adrian/gif-backend/utils"
)

func SearchGIFs(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    limit := r.URL.Query().Get("limit")
    page := r.URL.Query().Get("page")
    rating := r.URL.Query().Get("rating")

    if q == "" {
        http.Error(w, "Query param 'q' is required", http.StatusBadRequest)
        return
    }
    if limit == "" {
        limit = "12"
    }
    if page == "" {
        page = "1"
    }

    limitInt, _ := strconv.Atoi(limit)
    pageInt, _ := strconv.Atoi(page)

    if os.Getenv("GIPHY_API_KEY") == "" {
        http.Error(w, "Server misconfigured: missing GIPHY_API_KEY", http.StatusInternalServerError)
        return
    }

    // Typed call into our utils package
    result, err := utils.SearchGIFs(q, rating, limitInt, pageInt)
    if err != nil {
        http.Error(w, "Failed to fetch search results", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}
