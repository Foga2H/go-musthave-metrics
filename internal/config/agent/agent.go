package agent

import (
	"time"
)

type Config struct {
	PollInterval, ReportInterval time.Duration
}

func NewConfig() *Config {
	return &Config{
		PollInterval:   time.Second * 2,
		ReportInterval: time.Second * 10,
	}
}
