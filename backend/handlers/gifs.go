package handlers

import (
    "encoding/json"
    "net/http"
    "os"
    "strconv"

    "github.com/adrian/gif-backend/utils"
)

// ... (keep your corsMiddleware, imports, etc.)

func GetTrending(w http.ResponseWriter, r *http.Request) {
    // parse limit & page as before...
    limit := r.URL.Query().Get("limit")
    page := r.URL.Query().Get("page")
    if limit == "" { limit = "12" }
    if page == ""  { page  = "1" }

    limitInt, _ := strconv.Atoi(limit)
    pageInt,  _ := strconv.Atoi(page)
    
    // ensure API key is set
    if os.Getenv("GIPHY_API_KEY") == "" {
        http.Error(w, "Server misconfigured: missing GIPHY_API_KEY", http.StatusInternalServerError)
        return
    }

    // **Typed call**:
    respData, err := utils.FetchTrending(limitInt, pageInt)
    if err != nil {
        http.Error(w, "Failed to fetch trending GIFs", http.StatusInternalServerError)
        return
    }

    // JSON-encode the typed struct
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(respData)
}
