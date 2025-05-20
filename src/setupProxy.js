// Import the factory function to create a proxy middleware
const { createProxyMiddleware } = require("http-proxy-middleware");

/**
 * This function is automatically called by Create React App (when setupProxy.js is present).
 * It lets us proxy certain requests (e.g. API calls) to another server during development,
 * avoiding CORS issues and keeping frontend code free of absolute backend URLs.
 */
module.exports = function (app) {
  // Intercept any requests that begin with `/api`
  app.use(
    "/api",
    createProxyMiddleware({
      target: "http://localhost:5050", // the backend server to forward requests to  // The real backend address where API endpoints live.
      changeOrigin: true,              // modifies the request’s "Host" header to match the target
    })
  );
};
