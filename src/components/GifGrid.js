import React from "react";

function GifGrid({ gifs, favorites, toggleFavorite, copyLink }) {
  // 👇 FORCE an error to test ErrorBoundary
  if (process.env.NODE_ENV !== 'production') {
  throw new Error('💥 Test crash in GifGrid');
 }

  return (
    <div className="gif-grid">
      {gifs.map((gif) => (
        <div key={gif.id} className="gif-item">
          <div className="gif-image-wrapper">
            <img src={gif.images.fixed_height.url} alt={gif.title} />
            <div className="gif-buttons-overlay">
              <button
                className="top-left tooltip tooltip-bottom"
                onClick={() => toggleFavorite(gif)}
                data-tooltip={
                  favorites.some((f) => f.id === gif.id)
                    ? "Remove from favorites"
                    : "Add to favorites"
                }
              >
                {favorites.some((f) => f.id === gif.id) ? "💖" : "🤍"}
              </button>

              <button
                className="top-right tooltip tooltip-buttom"
                onClick={() => copyLink(gif.images.original.url)}
                data-tooltip="Copy link to clipboard"
              >
                📋
              </button>

              <a
                className="bottom-left tooltip tooltip-top"
                href={`https://wa.me/?text=${encodeURIComponent(
                  gif.images.original.url
                )}`}
                target="_blank"
                rel="noopener noreferrer"
                data-tooltip="Share on WhatsApp"
              >
                🟢
              </a>

              <a
                className="bottom-right tooltip tooltip-top"
                href={gif.images.original.url}
                download={`gif-${gif.id}.gif`}
                data-tooltip="Download GIF"
              >
                💾
              </a>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}

export default GifGrid;
