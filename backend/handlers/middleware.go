package handlers

import (
    "net/http"   // HTTP types for handlers
    "time"       // for measuring request duration

    // Prometheus client libraries for metrics
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    // Logrus for structured logging
    "github.com/sirupsen/logrus"
)

// Define Prometheus metrics as package-level variables so they’re
// shared across all requests.

// httpRequestsTotal counts the number of HTTP requests received,
// labeled by HTTP method, path, and response status.
var httpRequestsTotal = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "http_requests_total",
        Help: "Count of HTTP requests",
    },
    []string{"method", "path", "status"},
)

// httpRequestDuration tracks the duration of HTTP requests in seconds,
// labeled by HTTP method and path.
var httpRequestDuration = prometheus.NewHistogramVec(
    prometheus.HistogramOpts{
        Name:    "http_request_duration_seconds",
        Help:    "Duration of HTTP requests",
        Buckets: prometheus.DefBuckets, // default buckets provided by Prometheus
    },
    []string{"method", "path"},
)

// init registers our custom metrics with Prometheus’s default registry.
// Must happen before metrics can be collected.
func init() {
    prometheus.MustRegister(httpRequestsTotal, httpRequestDuration)
}

// LoggingAndMetricsMiddleware wraps each HTTP request to:
// 1) Record start time
// 2) Let the next handler run (capturing status via wrapper)
// 3) Log structured info via Logrus
// 4) Increment Prometheus counters and observe durations
func LoggingAndMetricsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 1) Note when the request started
        start := time.Now()

        // 2) Wrap ResponseWriter to capture the HTTP status code written
        lrw := &loggingResponseWriter{
            ResponseWriter: w,
            statusCode:     http.StatusOK, // default to 200 OK
        }

        // 3) Call the next handler in the chain
        next.ServeHTTP(lrw, r)

        // 4) After handler finishes, compute request duration
        duration := time.Since(start).Seconds()

        // 5) Extract method, path, and captured status
        method, path, status := r.Method, r.URL.Path, lrw.statusCode

        // 6) Log structured entry via Logrus
        logrus.WithFields(logrus.Fields{
            "method":   method,
            "path":     path,
            "status":   status,
            "duration": duration,
        }).Info("handled request")

        // 7) Update Prometheus metrics
        httpRequestsTotal.WithLabelValues(method, path, http.StatusText(status)).Inc()
        httpRequestDuration.WithLabelValues(method, path).Observe(duration)
    })
}

// loggingResponseWriter embeds http.ResponseWriter and intercepts
// WriteHeader calls to record the status code.
type loggingResponseWriter struct {
    http.ResponseWriter      // the original writer
    statusCode         int   // captured status code
}

// Override WriteHeader to capture the status code before sending it
func (lrw *loggingResponseWriter) WriteHeader(code int) {
    lrw.statusCode = code                  // record the code
    lrw.ResponseWriter.WriteHeader(code)  // forward the call
}

// ExposeMetricsHandler returns an HTTP handler that serves the
// /metrics endpoint for Prometheus to scrape.
func ExposeMetricsHandler() http.Handler {
    return promhttp.Handler()
}
