package utils

// GifImage holds a single image URL for a specific GIF variant.
// The `json:"url"` tag tells Go’s JSON decoder to fill this field from the “url” key.
type GifImage struct {
    URL string `json:"url"`
}

// Images bundles all the different size/format variants for a GIF.
// We only care about the “fixed_height” variant in our UI, but you could
// add others (e.g., original, downsized) as additional fields here.
type Images struct {
    // FixedHeight contains the URL and metadata for the fixed-height version.
    FixedHeight GifImage `json:"fixed_height"`
}

// Gif represents one GIF item returned by the Giphy API.
// It includes an ID, a human-readable title, and a collection of image variants.
type Gif struct {
    // ID is the unique identifier for this GIF.
    ID string `json:"id"`

    // Title is the descriptive text for the GIF.
    Title string `json:"title"`

    // Images holds different size variants; we use Images.FixedHeight in our grid.
    Images Images `json:"images"`
}

// Pagination contains metadata about the result set, such as total count,
// how many items are in this page, and the offset used for pagination.
type Pagination struct {
    // TotalCount is the total number of GIFs matching the query.
    TotalCount int `json:"total_count"`

    // Count is the number of GIFs returned in this response.
    Count int `json:"count"`

    // Offset is the zero-based index of the first GIF in this page.
    Offset int `json:"offset"`
}

// GiphyResponse mirrors the top-level JSON object returned by Giphy’s API.
// It contains the list of GIFs under Data and pagination info under Pagination.
type GiphyResponse struct {
    // Data is the slice of GIF items.
    Data []Gif `json:"data"`

    // Pagination holds paging metadata for the Data slice.
    Pagination Pagination `json:"pagination"`
}
