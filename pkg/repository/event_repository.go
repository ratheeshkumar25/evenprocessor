package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ratheeshkumar/event-processor/pkg/domain"
)

type eventRepository struct {
	webhookURL string
	client     *http.Client
}

func NewEventRepository(webhookURL string) domain.EventRepository {
	return &eventRepository{
		webhookURL: webhookURL,
		client:     &http.Client{},
	}
}

func (r *eventRepository) SendToWebhook(event *domain.Event) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	req, err := http.NewRequest("POST", r.webhookURL, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("webhook returned error status: %d", resp.StatusCode)
	}

	return nil
}
