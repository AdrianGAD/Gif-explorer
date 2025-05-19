package utils

// GifImage holds a single image URL.
type GifImage struct {
    URL string `json:"url"`
}

// Images bundles all the different size variants.
// We care about fixed_height for our grid.
type Images struct {
    FixedHeight GifImage `json:"fixed_height"`
}

// Gif represents one GIF item from Giphy.
type Gif struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Images Images `json:"images"`
}

// Pagination metadata for result sets.
type Pagination struct {
    TotalCount int `json:"total_count"`
    Count      int `json:"count"`
    Offset     int `json:"offset"`
}

// GiphyResponse mirrors the top‐level object returned by Giphy.
type GiphyResponse struct {
    Data       []Gif      `json:"data"`
    Pagination Pagination `json:"pagination"`
}
