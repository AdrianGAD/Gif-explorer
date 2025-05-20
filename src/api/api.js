// Base URL of the backend API. In production we might load this from an environment variable.
const BACKEND = "http://localhost:5050";

/**
 * searchGIFs
 *  - q:      the search query string (e.g., "cats")
 *  - page:   which page of results (for pagination), defaults to 1
 *  - rating: content rating filter (e.g., "g", "pg"), defaults to "g"
 *  - lang:   language for query parsing (e.g., "en"), defaults to "en"
 *  - limit:  number of GIFs to return per page, defaults to 12
 */
export async function searchGIFs(
  q,
  page = 1,
  rating = "g",
  lang = "en",
  limit = 12
) {
  // 1) Build the query parameters string, URL-escaping the search term.
  const params =
    `?q=${encodeURIComponent(q)}` +
    `&page=${page}` +
    `&limit=${limit}` +
    `&rating=${rating}` +
    `&lang=${lang}`;

  // 2) Combine the backend base URL, API path, and query params.
  const url = `${BACKEND}/api/search${params}`;

  // 3) Perform the HTTP GET request to our Go backend.
  const res = await fetch(url);

  // 4) If the response status is not in the 200–299 range, throw an error.
  if (!res.ok) {
    // This will be caught by the caller to show an error message in the UI.
    throw new Error("Failed to search GIFs");
  }

  // 5) Parse and return the JSON body. The shape matches our GiphyResponse struct
  //    transformed by the backend: { data: [...], pagination: {...} }
  return res.json();
}
