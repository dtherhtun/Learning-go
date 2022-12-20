package main

import (
	"embed"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	//go:embed html/index.html
	content embed.FS

	totalRequest = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "opslab_http_requests_total",
			Help: "Number of get requests.",
		},
		[]string{"path"},
	)

	responseStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "opslab_response_status",
			Help: "Status of opslab HTTP response",
		},
		[]string{"status"},
	)

	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "opslab_http_response_time_seconds",
			Help: "Duration of opslab HTTP requests",
		}, []string{"path"},
	)
	helloLast = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "opslab_last_time_seconds",
			Help: "The last time a hello World was served",
		},
	)
	opslabLatency = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Name: "opslab_latency_seconds",
			Help: "Time for a request Hello World.",
		})
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func init() {
	prometheus.Register(totalRequest)
	prometheus.Register(responseStatus)
	prometheus.Register(httpDuration)
	prometheus.Register(helloLast)
	prometheus.Register(opslabLatency)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/", mw(http.StripPrefix("/", http.FileServer(http.FS(content)))))
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err)
	}
}

func mw(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		start := time.Now()
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		defer timer.ObserveDuration()

		rw := NewResponseWriter(w)
		h.ServeHTTP(rw, r)

		statusCode := rw.statusCode
		responseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc()

		totalRequest.WithLabelValues(path).Inc()
		helloLast.SetToCurrentTime()
		opslabLatency.Observe(float64(time.Now().UnixNano() - start.UnixNano()))
	})
}
