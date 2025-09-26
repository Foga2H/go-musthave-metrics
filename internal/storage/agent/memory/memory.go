package agent

import "github.com/Foga2H/go-musthave-metrics/internal/storage/agent"

type MemStorage struct {
	metrics agent.Metrics
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		metrics: agent.Metrics{},
	}
}

func (m *MemStorage) Save(metrics agent.Metrics) {
	m.metrics = metrics
}

func (m *MemStorage) GetAll() agent.Metrics {
	return m.metrics
}
