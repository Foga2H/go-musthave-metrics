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
		err := sendRequest(
			value.Type,
			key,
			strconv.FormatFloat(value.Value, 'f', -1, 64),
		)

		if err != nil {
			return err
		}
	}

	fmt.Printf("Successfully updated metrics: %v\n", metrics.Data)

	return nil
}

func sendRequest(typeField string, nameField string, valueField string) error {
	response, err := http.Post(fmt.Sprintf(
		"http://localhost:8080/update/%s/%s/%s",
		typeField,
		nameField,
		valueField,
	), "text/plain", strings.NewReader(""))

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error updating metrics: %s\n", response.Status)
	}

	return nil
}
