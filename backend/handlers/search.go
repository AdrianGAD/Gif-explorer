package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// NOTE: no const GiphyBaseURL here — it’s defined once in gifs.go

func SearchGIFs(w http.ResponseWriter, r *http.Request) {
	// pull query params
	query := r.URL.Query().Get("q")
	limit := r.URL.Query().Get("limit")
	page := r.URL.Query().Get("page")
	rating := r.URL.Query().Get("rating")
	lang := r.URL.Query().Get("lang")

	if query == "" {
		http.Error(w, "Query param 'q' is required", http.StatusBadRequest)
		return
	}
	if limit == "" {
		limit = "12"
	}
	if page == "" {
		page = "1"
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 12
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	offset := (pageInt - 1) * limitInt

	// load API key now that .env has been loaded
	apiKey := os.Getenv("GIPHY_API_KEY")
	if apiKey == "" {
		http.Error(w, "Server misconfigured: missing GIPHY_API_KEY", http.StatusInternalServerError)
		return
	}

	// build and call the Giphy search endpoint
	endpoint := fmt.Sprintf(
		"%s/search?q=%s&limit=%d&offset=%d&api_key=%s",
		GiphyBaseURL, url.QueryEscape(query), limitInt, offset, apiKey,
	)
	if rating != "" {
		endpoint += "&rating=" + rating
	}
	if lang != "" {
		endpoint += "&lang=" + lang
	}

	resp, err := http.Get(endpoint)
	if err != nil {
		http.Error(w, "Failed to fetch search results", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// forward the JSON response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read GIFs response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
