import React from "react";
import GifGrid from "./GifGrid";

function FavoritesSection({ favorites, toggleFavorite, copyLink }) {
  if (favorites.length === 0) return null;

  return (
    <div className="favorites-section">
      <h2>Your Favorites</h2>
      <GifGrid
        gifs={favorites}
        favorites={favorites}
        toggleFavorite={toggleFavorite}
        copyLink={copyLink}
      />
    </div>
  );
}

export default FavoritesSection;
