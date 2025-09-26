package agent

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Foga2H/go-musthave-metrics/internal/storage/agent"
)

type SenderService struct{}

func NewSenderService() *SenderService {
	return &SenderService{}
}

func (s *SenderService) Send(metrics agent.Metrics) error {
	for key, value := range metrics.Data {
		response, err := http.Post(fmt.Sprintf(
			"http://localhost:8080/update/%s/%s/%s",
			value.Type,
			key,
			strconv.FormatFloat(value.Value, 'f', -1, 64),
		), "text/plain", strings.NewReader(""))

		if err != nil {
			panic(err)
		}

		if response.StatusCode != http.StatusOK {
			fmt.Printf("Error updating metrics: %s\n", response.Status)
		}
	}

	fmt.Printf("Successfully updated metrics: %s\n", metrics.Data)

	return nil
}
