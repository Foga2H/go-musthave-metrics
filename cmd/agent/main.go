package main

import (
	"fmt"
	"time"

	"github.com/Foga2H/go-musthave-metrics/internal/config/agent"
	agent2 "github.com/Foga2H/go-musthave-metrics/internal/service/agent"
	agent3 "github.com/Foga2H/go-musthave-metrics/internal/storage/agent/memory"
)

var pollCount int64 = 0

func main() {
	agentConfig := agent.NewConfig()
	collectService := agent2.NewCollectService()
	senderService := agent2.NewSenderService()
	agentStorage := agent3.NewMemStorage()

	go func() {
		for {
			time.Sleep(agentConfig.ReportInterval)
			metrics := agentStorage.GetAll()
			err := senderService.Send(metrics)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	for {
		pollCount = pollCount + 1
		fmt.Printf("Getting metrics. Poll count: %d\n", pollCount)
		metrics := collectService.Collect(pollCount)
		agentStorage.Save(metrics)

		time.Sleep(agentConfig.PollInterval)
	}
}
