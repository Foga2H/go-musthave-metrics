package main

import (
	"net/http"

	"github.com/Foga2H/go-musthave-metrics/internal/handler"
	"github.com/Foga2H/go-musthave-metrics/internal/storage/server/memory"
)

func main() {
	mux := http.NewServeMux()
	storage := server.NewMemStorage()

	mux.Handle(`/update/{type}/{metric}/{value}`, &handler.MetricsHandler{
		Storage: storage,
	})
	mux.Handle(`/`, &handler.NotFoundHandler{})

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
