const BACKEND = "http://localhost:5050"; // or import from env var

export async function searchGIFs(
  q,
  page = 1,
  rating = "g",
  lang = "en",
  limit = 12
) {
  const url =
    `${BACKEND}/api/search` +
    `?q=${encodeURIComponent(q)}` +
    `&page=${page}&limit=${limit}&rating=${rating}&lang=${lang}`;
  const res = await fetch(url);
  if (!res.ok) throw new Error("Failed to search GIFs");
  return res.json();
}
