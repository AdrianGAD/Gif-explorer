import React from "react";

export default class ErrorBoundary extends React.Component {
  constructor(props) {
    super(props);
    this.state = { hasError: false };
  }

  static getDerivedStateFromError(error) {
    // Update state so the next render shows the fallback UI.
    return { hasError: true };
  }

  componentDidCatch(error, info) {
    // You can log error details to an external service here.
    console.error("ErrorBoundary caught an error:", error, info);
  }

  render() {
    if (this.state.hasError) {
      // Fallback UI
      return (
        <div style={{ padding: "2rem", textAlign: "center" }}>
          <h2>Oops — something went wrong.</h2>
          <p>Please refresh the page or try again later.</p>
        </div>
      );
    }

    // If no error, render children as normal
    return this.props.children;
  }
}
