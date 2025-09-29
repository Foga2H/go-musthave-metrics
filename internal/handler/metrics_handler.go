package handler

import (
	"fmt"
	"net/http"
	"strconv"

	metrics "github.com/Foga2H/go-musthave-metrics/internal/model"
	"github.com/Foga2H/go-musthave-metrics/internal/repository/server"
)

type MetricsHandler struct {
	Storage server.StorageRepo
}

func NewMetricsHandler(storage server.StorageRepo) *MetricsHandler {
	return &MetricsHandler{
		Storage: storage,
	}
}

func (h *MetricsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("content-type", "text/plain")

	typeField := r.PathValue("type")
	metricField := r.PathValue("metric")
	valueField := r.PathValue("value")

	if metricField == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if typeField != metrics.Counter && typeField != metrics.Gauge {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if typeField == metrics.Counter {
		if _, err := strconv.ParseInt(valueField, 10, 64); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if typeField == metrics.Gauge {
		if _, err := strconv.ParseFloat(valueField, 64); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	h.Storage.Set(typeField, valueField)
	fmt.Printf("Got metrics for type: %s, metric: %s, value: %s\n", typeField, metricField, valueField)

	w.WriteHeader(http.StatusOK)
}

func notFound(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusNotFound)
}
