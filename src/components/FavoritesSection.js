import React from "react";
import GifGrid from "./GifGrid";

/**
 * FavoritesSection displays a list of the user’s favorite GIFs.
 *
 * Props:
 * - favorites: Array of GIF objects that the user has marked as favorites.
 * - toggleFavorite: Function to add/remove a GIF from favorites.
 * - copyLink: Function to copy a GIF’s URL to the clipboard.
 */
function FavoritesSection({ favorites, toggleFavorite, copyLink }) {
  // If there are no favorites, render nothing (null) to avoid empty UI.
  if (favorites.length === 0) return null;

  // Otherwise, render a section containing a heading and the GifGrid.
  return (
    <div className="favorites-section">
      {/* Section title */}
      <h2>Your Favorites</h2>

      {/* Reuse GifGrid to display the list of favorite GIFs.
          We pass:
          - gifs: the array of favorite GIFs to render
          - favorites: same array so GifGrid can tell which items are “favorited”
          - toggleFavorite: handler to add/remove a GIF when its heart button is clicked
          - copyLink: handler to copy the original GIF URL when the copy button is clicked
      */}
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
