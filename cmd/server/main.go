package main

import (
	"net/http"
	"strconv"

	metrics "github.com/Foga2H/go-musthave-metrics/internal/model"
	storageInterface "github.com/Foga2H/go-musthave-metrics/internal/storage"
	"github.com/Foga2H/go-musthave-metrics/internal/storage/memory"
)

var storage storageInterface.Storage

func updateMetrics(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	res.Header().Set("content-type", "text/plain")

	typeField := req.PathValue("type")
	metricField := req.PathValue("metric")
	valueField := req.PathValue("value")

	if metricField == "" {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	if typeField != metrics.Counter && typeField != metrics.Gauge {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	if typeField == metrics.Counter {
		if _, err := strconv.ParseInt(valueField, 10, 64); err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if typeField == metrics.Gauge {
		if _, err := strconv.ParseFloat(valueField, 64); err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	storage.Set(typeField, valueField)

	res.WriteHeader(http.StatusOK)
}

func notFound(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusNotFound)
}

func main() {
	mux := http.NewServeMux()
	storage = memory.NewMemStorage()

	mux.HandleFunc(`/update/{type}/{metric}/{value}`, updateMetrics)
	mux.HandleFunc(`/`, notFound)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
