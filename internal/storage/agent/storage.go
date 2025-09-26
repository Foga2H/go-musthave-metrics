package agent

type Metric struct {
	Type  string
	Value float64
}

type Metrics struct {
	Data map[string]Metric
}

type Storage struct {
	Metrics Metrics
}
