import React from "react";

/**
 * ErrorBoundary is a React component that catches JavaScript errors
 * anywhere in its child component tree, logs those errors, and displays
 * a fallback UI instead of the component tree that crashed.
 */
export default class ErrorBoundary extends React.Component {
  constructor(props) {
    super(props);
    // Initialize state to track whether an error has been caught.
    this.state = { hasError: false };
  }

  /**
   * getDerivedStateFromError is a React lifecycle method.
   * It runs when a child component throws an error during rendering,
   * in a lifecycle method, or in the constructor of any child.
   * Returning a new state object here updates the state so that
   * render() shows the fallback UI.
   */
  static getDerivedStateFromError(error) {
    // We don't use the error object directly here; we just flip the flag.
    return { hasError: true };
  }

  /**
   * componentDidCatch is another React lifecycle method.
   * It’s called after an error has been thrown by a child component.
   * You can use it to log error details or send them to an external service.
   */
  componentDidCatch(error, info) {
    // Example: log the error and component stack trace to the console.
    console.error("ErrorBoundary caught an error:", error, info);
    // In production, we might send 'error' and 'info.componentStack'
    // to a monitoring service here.
  }

  render() {
    // If an error was caught, display a user-friendly fallback UI.
    if (this.state.hasError) {
      return (
        <div style={{ padding: "2rem", textAlign: "center" }}>
          <h2>Oops — something went wrong.</h2>
          <p>Please refresh the page or try again later.</p>
        </div>
      );
    }

    // If no error occurred, render children components normally.
    return this.props.children;
  }
}
