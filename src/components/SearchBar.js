import React from "react";

/**
 * SearchBar renders an input field and a button for searching GIFs.
 *
 * Props:
 * - searchTerm: current text in the search input
 * - setSearchTerm: function to update the searchTerm state in the parent
 * - handleSearch: function to trigger the search action (e.g., fetch GIFs)
 */
function SearchBar({ searchTerm, setSearchTerm, handleSearch }) {
  return (
    <div className="search-bar">
      {/*
        The text input is controlled by the `searchTerm` prop.
        onChange updates the parent state via setSearchTerm.
        onKeyDown listens for the Enter key to trigger a search.
      */}
      <input
        type="text"
        placeholder="Search for GIFs"
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
        onKeyDown={(e) => e.key === "Enter" && handleSearch()}
      />
      {/*
        The Search button invokes handleSearch when clicked.
      */}
      <button onClick={handleSearch}>Search</button>
    </div>
  );
}

export default SearchBar;
