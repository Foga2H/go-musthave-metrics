package server

type MemStorage struct {
	metrics map[string]string
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		metrics: make(map[string]string),
	}
}

func (m *MemStorage) Set(key string, value string) {
	m.metrics[key] = value
}

func (m *MemStorage) Get(key string) (string, bool) {
	val, ok := m.metrics[key]
	return val, ok
}
