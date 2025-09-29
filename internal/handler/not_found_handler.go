package handler

import (
	"net/http"

	"github.com/Foga2H/go-musthave-metrics/internal/repository/server"
)

type NotFoundHandler struct{}

func NewNotFoundHandler(storage server.StorageRepo) *NotFoundHandler {
	return &NotFoundHandler{}
}

func (h *NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusNotFound)
}
