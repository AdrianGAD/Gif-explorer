import React, { useEffect, useState, useCallback } from "react";
import { searchGIFs } from "./api/api";
import "./styles/App.css";

function App() {
  // ─────────────────────────────────────────────────────────────────────────────
  // State variables
  // ─────────────────────────────────────────────────────────────────────────────
  // gifs: array of GIF objects to display
  const [gifs, setGifs] = useState([]);
  // query: current search term
  const [query, setQuery] = useState("");
  // favorites: list of user-favorited GIFs
  const [favorites, setFavorites] = useState([]);
  // activeTab: which view is shown ("trending" or "favorites")
  const [activeTab, setActiveTab] = useState("trending");
  // loading: whether a fetch is in progress
  const [loading, setLoading] = useState(false);
  // page: current pagination page for trending/search
  const [page, setPage] = useState(1);
  // animateId: temporarily holds a GIF id to trigger a favorite animation
  const [animateId, setAnimateId] = useState(null);
  // rating & lang: filters for search requests, persisted to localStorage
  const [rating, setRating] = useState("g");
  const [lang, setLang] = useState("en");

  // Base URL for our backend API
  const BACKEND_URL = "http://localhost:5050/api";

  // ─────────────────────────────────────────────────────────────────────────────
  // Persisted filters in localStorage
  // ─────────────────────────────────────────────────────────────────────────────
  // On mount, load saved rating & lang
  useEffect(() => {
    const savedRating = localStorage.getItem("rating");
    const savedLang    = localStorage.getItem("lang");
    if (savedRating) setRating(savedRating);
    if (savedLang)    setLang(savedLang);
  }, []);

  // Whenever rating changes, persist it
  useEffect(() => {
    localStorage.setItem("rating", rating);
  }, [rating]);

  // Whenever lang changes, persist it
  useEffect(() => {
    localStorage.setItem("lang", lang);
  }, [lang]);

  // ─────────────────────────────────────────────────────────────────────────────
  // Fetch Trending GIFs
  // ─────────────────────────────────────────────────────────────────────────────
  // Wrap in useCallback so it only changes when `page` changes
  const fetchTrending = useCallback(async () => {
    setLoading(true);
    try {
      // Call our backend directly
      const res  = await fetch(`${BACKEND_URL}/trending?limit=12&page=${page}`);
      const data = await res.json();
      // GiphyResponse has shape { data: [...], pagination: {...} }
      setGifs(data.data);
    } catch (err) {
      console.error("Failed to fetch trending GIFs", err);
    } finally {
      setLoading(false);
    }
  }, [page]);

  // When activeTab switches to "trending" or page changes, re-fetch
  useEffect(() => {
    if (activeTab === "trending") {
      fetchTrending();
    }
  }, [activeTab, fetchTrending]);

  // ─────────────────────────────────────────────────────────────────────────────
  // Search handler
  // ─────────────────────────────────────────────────────────────────────────────
  const handleSearch = async () => {
    setLoading(true);
    try {
      // Use our API helper for search, passing query, page, rating, lang
      const results = await searchGIFs(query, page, rating, lang);
      setGifs(results.data);
    } catch (err) {
      console.error("Search failed", err);
    } finally {
      setLoading(false);
    }
  };

  // ─────────────────────────────────────────────────────────────────────────────
  // Favorites management
  // ─────────────────────────────────────────────────────────────────────────────
  const toggleFavorite = (gif) => {
    const exists = favorites.some((f) => f.id === gif.id);
    if (exists) {
      // Remove from favorites
      setFavorites(favorites.filter((f) => f.id !== gif.id));
    } else {
      // Add to favorites and trigger animation
      setFavorites([...favorites, gif]);
      setAnimateId(gif.id);
      setTimeout(() => setAnimateId(null), 300);
    }
  };
  const isFavorited = (gif) => favorites.some((f) => f.id === gif.id);

  // ─────────────────────────────────────────────────────────────────────────────
  // Pagination controls
  // ─────────────────────────────────────────────────────────────────────────────
  const handlePrev = () => setPage((p) => Math.max(1, p - 1));
  const handleNext = () => setPage((p) => p + 1);

  // ─────────────────────────────────────────────────────────────────────────────
  // JSX layout
  // ─────────────────────────────────────────────────────────────────────────────
  return (
    <div className="app">
      {/* Header & Tab Navigation */}
      <header className="app-header">
        <h1>GIF Explorer</h1>
        <div className="nav">
          <span
            className={`nav-item ${activeTab === "trending" ? "active" : ""}`}
            onClick={() => {
              setActiveTab("trending");
              setPage(1);
            }}
          >
            🔥 Trending
          </span>
          <span
            className={`nav-item ${activeTab === "favorites" ? "active" : ""}`}
            onClick={() => setActiveTab("favorites")}
          >
            💖 Favorites
          </span>
        </div>
      </header>

      {/* Filters & Search Bar (only in Trending tab) */}
      {activeTab === "trending" && (
        <>
          <div className="filters">
            <label title="Filter by content rating">
              Rating:
              <select
                value={rating}
                onChange={(e) => setRating(e.target.value)}
              >
                <option value="g">G</option>
                <option value="pg">PG</option>
                <option value="pg-13">PG-13</option>
                <option value="r">R</option>
              </select>
            </label>
            <label title="Filter by language for search accuracy">
              Language:
              <select value={lang} onChange={(e) => setLang(e.target.value)}>
                <option value="en">English</option>
                <option value="es">Spanish</option>
                <option value="fr">French</option>
                <option value="pt">Portuguese</option>
                <option value="de">German</option>
              </select>
            </label>
            <button
              className="reset-btn"
              onClick={() => {
                // Reset filters back to defaults
                setRating("g");
                setLang("en");
                localStorage.setItem("rating", "g");
                localStorage.setItem("lang", "en");
              }}
            >
              Reset Filters
            </button>
          </div>

          {/* Search input */}
          <div className="search-bar">
            <input
              type="text"
              placeholder="Search for GIFs"
              value={query}
              onChange={(e) => setQuery(e.target.value)}
              onKeyDown={(e) => e.key === "Enter" && handleSearch()}
            />
            <button onClick={handleSearch}>Search</button>
          </div>
        </>
      )}

      {/* Loading spinner */}
      {loading ? (
        <div className="spinner">Loading...</div>
      ) : (
        /* GIF Grid or Favorites Grid */
        <div className="gif-grid">
          {(activeTab === "trending" ? gifs : favorites).map((gif) => (
            <div className="gif-card" key={gif.id}>
              <div className="gif-wrapper">
                <img src={gif.images.fixed_height.url} alt={gif.title} />
                <button
                  className={`heart-icon ${
                    isFavorited(gif) ? "active" : ""
                  } ${animateId === gif.id ? "animate" : ""}`}
                  onClick={() => toggleFavorite(gif)}
                >
                  ❤️
                </button>
              </div>
              <p className="gif-title">{gif.title}</p>
            </div>
          ))}
        </div>
      )}

      {/* No favorites message */}
      {activeTab === "favorites" && favorites.length === 0 && !loading && (
        <p>😔 No favorite GIFs yet.</p>
      )}

      {/* Pagination controls (only in Trending tab when gifs exist) */}
      {activeTab === "trending" && !loading && gifs.length > 0 && (
        <div className="pagination">
          <button onClick={handlePrev} disabled={page === 1}>
            Previous
          </button>
          <span>Page {page}</span>
          <button onClick={handleNext}>Next</button>
        </div>
      )}
    </div>
  );
}

export default App;
