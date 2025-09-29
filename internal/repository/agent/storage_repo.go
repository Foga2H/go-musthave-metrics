package agent

import "github.com/Foga2H/go-musthave-metrics/internal/storage/agent"

type StorageRepo interface {
	Save(metrics agent.Metrics)
	GetAll() agent.Metrics
}
