package utils

import (
    "encoding/json" // for decoding JSON responses
    "fmt"           // for building URLs with fmt.Sprintf
    "net/http"      // for making HTTP requests
    "os"            // for reading environment variables
)

// baseURL is the Giphy API endpoint root. 
// We declare it as a var so tests can override it if needed.
var baseURL = "https://api.giphy.com/v1/gifs"

// FetchTrending retrieves the current trending GIFs from Giphy.
// It returns a typed GiphyResponse or an error.
func FetchTrending(limit, page int) (GiphyResponse, error) {
    // 1) Calculate pagination offset.
    offset := (page - 1) * limit

    // 2) Build the request URL including API key, limit, offset, and a 'g' rating filter.
    url := fmt.Sprintf(
        "%s/trending?api_key=%s&limit=%d&offset=%d&rating=g",
        baseURL,
        os.Getenv("GIPHY_API_KEY"), // read API key at runtime
        limit,
        offset,
    )

    // 3) Perform the HTTP GET request.
    resp, err := http.Get(url)
    if err != nil {
        // Network or DNS error: bubble up to caller.
        return GiphyResponse{}, err
    }
    // Ensure the response body is closed after decoding.
    defer resp.Body.Close()

    // 4) Decode the JSON response into our typed struct.
    var result GiphyResponse
    err = json.NewDecoder(resp.Body).Decode(&result)
    // Return the decoded struct (even if err is non-nil, so caller sees partial data).
    return result, err
}

// SearchGIFs queries Giphy for GIFs matching the given search term.
// It accepts a rating filter, pagination parameters, and returns a GiphyResponse.
func SearchGIFs(query, rating string, limit, page int) (GiphyResponse, error) {
    // 1) Calculate pagination offset.
    offset := (page - 1) * limit

    // 2) Build the search URL with API key, escaped query, limit, offset, and rating.
    url := fmt.Sprintf(
        "%s/search?api_key=%s&q=%s&limit=%d&offset=%d&rating=%s",
        baseURL,
        os.Getenv("GIPHY_API_KEY"), // fetch API key
        query,                      // search term (assumed already URL-escaped by caller if needed)
        limit,
        offset,
        rating,                     // content rating filter
    )

    // 3) Perform the HTTP GET request against Giphy’s search endpoint.
    resp, err := http.Get(url)
    if err != nil {
        // Return zero-value struct and the error.
        return GiphyResponse{}, err
    }
    defer resp.Body.Close()

    // 4) Decode JSON into our typed response struct.
    var result GiphyResponse
    err = json.NewDecoder(resp.Body).Decode(&result)
    return result, err
}
