package agent

import (
	"math/rand/v2"
	"runtime"

	metrics "github.com/Foga2H/go-musthave-metrics/internal/model"
	"github.com/Foga2H/go-musthave-metrics/internal/storage/agent"
)

type CollectService struct{}

func NewCollectService() *CollectService {
	return &CollectService{}
}

func (m *CollectService) Collect(pollCount int64) agent.Metrics {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	randomFloat := rand.Float64()

	data := map[string]agent.Metric{
		"Alloc": {
			Type:  metrics.Gauge,
			Value: float64(memStats.Alloc),
		},
		"BuckHashSys": {
			Type:  metrics.Gauge,
			Value: float64(memStats.BuckHashSys),
		},
		"Frees": {
			Type:  metrics.Gauge,
			Value: float64(memStats.Frees),
		},
		"GCCPUFraction": {
			Type:  metrics.Gauge,
			Value: memStats.GCCPUFraction,
		},
		"GCSys": {
			Type:  metrics.Gauge,
			Value: float64(memStats.GCSys),
		},
		"HeapAlloc": {
			Type:  metrics.Gauge,
			Value: float64(memStats.HeapAlloc),
		},
		"HeapIdle": {
			Type:  metrics.Gauge,
			Value: float64(memStats.HeapIdle),
		},
		"HeapInuse": {
			Type:  metrics.Gauge,
			Value: float64(memStats.HeapInuse),
		},
		"HeapObjects": {
			Type:  metrics.Gauge,
			Value: float64(memStats.HeapObjects),
		},
		"HeapReleased": {
			Type:  metrics.Gauge,
			Value: float64(memStats.HeapReleased),
		},
		"HeapSys": {
			Type:  metrics.Gauge,
			Value: float64(memStats.HeapSys),
		},
		"LastGC": {
			Type:  metrics.Gauge,
			Value: float64(memStats.LastGC),
		},
		"Lookups": {
			Type:  metrics.Gauge,
			Value: float64(memStats.Lookups),
		},
		"MCacheInuse": {
			Type:  metrics.Gauge,
			Value: float64(memStats.MCacheInuse),
		},
		"MCacheSys": {
			Type:  metrics.Gauge,
			Value: float64(memStats.MCacheSys),
		},
		"MSpanInuse": {
			Type:  metrics.Gauge,
			Value: float64(memStats.MSpanInuse),
		},
		"MSpanSys": {
			Type:  metrics.Gauge,
			Value: float64(memStats.MSpanSys),
		},
		"Mallocs": {
			Type:  metrics.Gauge,
			Value: float64(memStats.Mallocs),
		},
		"NextGC": {
			Type:  metrics.Gauge,
			Value: float64(memStats.NextGC),
		},
		"NumForcedGC": {
			Type:  metrics.Gauge,
			Value: float64(memStats.NumForcedGC),
		},
		"NumGC": {
			Type:  metrics.Gauge,
			Value: float64(memStats.NumGC),
		},
		"OtherSys": {
			Type:  metrics.Gauge,
			Value: float64(memStats.OtherSys),
		},
		"PauseTotalNs": {
			Type:  metrics.Gauge,
			Value: float64(memStats.PauseTotalNs),
		},
		"StackInuse": {
			Type:  metrics.Gauge,
			Value: float64(memStats.StackInuse),
		},
		"StackSys": {
			Type:  metrics.Gauge,
			Value: float64(memStats.StackSys),
		},
		"Sys": {
			Type:  metrics.Gauge,
			Value: float64(memStats.Sys),
		},
		"TotalAlloc": {
			Type:  metrics.Gauge,
			Value: float64(memStats.TotalAlloc),
		},
		"RandomValue": {
			Type:  metrics.Gauge,
			Value: randomFloat,
		},
		"PollCount": {
			Type:  metrics.Counter,
			Value: float64(pollCount),
		},
	}

	return agent.Metrics{
		Data: data,
	}
}
