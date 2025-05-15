package handlers

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

// Prom metrics
var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Count of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal, httpRequestDuration)
}

// LoggingAndMetricsMiddleware wraps handlers to log and instrument
func LoggingAndMetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// wrap ResponseWriter to capture status code
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(lrw, r)

		duration := time.Since(start).Seconds()
		method, path, status := r.Method, r.URL.Path, lrw.statusCode

		// Logrus structured log
		logrus.WithFields(logrus.Fields{
			"method":   method,
			"path":     path,
			"status":   status,
			"duration": duration,
		}).Info("handled request")

		// Prometheus metrics
		httpRequestsTotal.WithLabelValues(method, path, http.StatusText(status)).Inc()
		httpRequestDuration.WithLabelValues(method, path).Observe(duration)
	})
}

// helper to capture status code
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// ExposeMetricsHandler returns the promhttp handler
func ExposeMetricsHandler() http.Handler {
	return promhttp.Handler()
}
