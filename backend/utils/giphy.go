package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const baseURL = "https://api.giphy.com/v1/gifs"

func FetchTrending(limit, page int) (map[string]interface{}, error) {
	offset := (page - 1) * limit
	url := fmt.Sprintf("%s/trending?api_key=%s&limit=%d&offset=%d&rating=g",
		baseURL,
		os.Getenv("GIPHY_API_KEY"),
		limit,
		offset,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

func SearchGIFs(query, rating string, limit, page int) (map[string]interface{}, error) {
	offset := (page - 1) * limit
	url := fmt.Sprintf("%s/search?api_key=%s&q=%s&limit=%d&offset=%d&rating=%s",
		baseURL,
		os.Getenv("GIPHY_API_KEY"),
		query,
		limit,
		offset,
		rating,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]any
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}
