import React, { useEffect, useState, useCallback} from "react";
import { searchGIFs } from "./api/api";
import "./styles/App.css";

function App() {
  const [gifs, setGifs] = useState([]);
  const [query, setQuery] = useState("");
  const [favorites, setFavorites] = useState([]);
  const [activeTab, setActiveTab] = useState("trending");
  const [loading, setLoading] = useState(false);
  const [page, setPage] = useState(1);
  const [animateId, setAnimateId] = useState(null);
  const [rating, setRating] = useState("g");
  const [lang, setLang] = useState("en");

  const BACKEND_URL = "http://localhost:5050/api";

  useEffect(() => {
    const savedRating = localStorage.getItem("rating");
    const savedLang = localStorage.getItem("lang");
    if (savedRating) setRating(savedRating);
    if (savedLang) setLang(savedLang);
  }, []);

  useEffect(() => {
    localStorage.setItem("rating", rating);
  }, [rating]);

  useEffect(() => {
    localStorage.setItem("lang", lang);
  }, [lang]);

  const fetchTrending = useCallback(async () => {
  setLoading(true);
  try {
    const res = await fetch(`${BACKEND_URL}/trending?limit=12&page=${page}`);
    const data = await res.json();
    setGifs(data.data);
  } catch (err) {
    console.error("Failed to fetch trending GIFs", err);
  } finally {
    setLoading(false);
  }
}, [page]);

useEffect(() => {
  if (activeTab === "trending") {
    fetchTrending();
  }
}, [activeTab, fetchTrending]);

  const handleSearch = async () => {
    try {
      setLoading(true);
      const results = await searchGIFs(query, page, rating, lang);
      setGifs(results.data);
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const toggleFavorite = (gif) => {
    const exists = favorites.find((f) => f.id === gif.id);
    if (exists) {
      setFavorites(favorites.filter((f) => f.id !== gif.id));
    } else {
      setFavorites([...favorites, gif]);
      setAnimateId(gif.id);
      setTimeout(() => setAnimateId(null), 300);
    }
  };

  const isFavorited = (gif) => favorites.some((f) => f.id === gif.id);

  const handlePrev = () => setPage((p) => Math.max(1, p - 1));
  const handleNext = () => setPage((p) => p + 1);

  return (
    <div className="app">
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

      {activeTab === "trending" && (
        <>
          <div className="filters">
            <label title="Filter by content rating (G = general, R = adult)">
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
            <label title="Filter by GIF metadata language (affects search accuracy)">
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
                setRating("g");
                setLang("en");
                localStorage.setItem("rating", "g");
                localStorage.setItem("lang", "en");
              }}
            >
              Reset Filters
            </button>
          </div>

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

      {loading ? (
        <div className="spinner">Loading...</div>
      ) : (
        <div className="gif-grid">
          {(activeTab === "trending" ? gifs : favorites).map((gif) => (
            <div className="gif-card" key={gif.id}>
              <div className="gif-wrapper">
                <img src={gif.images.fixed_height.url} alt={gif.title} />
                <button
                  className={`heart-icon ${isFavorited(gif) ? "active" : ""} ${
                    animateId === gif.id ? "animate" : ""
                  }`}
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

      {activeTab === "favorites" && favorites.length === 0 && !loading && (
        <p>😔 No favorite GIFs yet.</p>
      )}

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
