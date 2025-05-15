package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

const GiphyBaseURL = "https://api.giphy.com/v1/gifs"

func GetTrending(w http.ResponseWriter, r *http.Request) {
	// parse pagination params
	limit := r.URL.Query().Get("limit")
	page := r.URL.Query().Get("page")
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

	// build and call the Giphy endpoint
	endpoint := fmt.Sprintf(
		"%s/trending?limit=%d&offset=%d&api_key=%s",
		GiphyBaseURL, limitInt, offset, apiKey,
	)
	resp, err := http.Get(endpoint)
	if err != nil {
		http.Error(w, "Failed to fetch trending GIFs", http.StatusInternalServerError)
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
