package usecase

import (
	"fmt"
	"log"
	"strings"

	"github.com/ratheeshkumar/event-processor/pkg/domain"
)

type eventUseCase struct {
	repo domain.EventRepository
}

func NewEventUseCase(repo domain.EventRepository) domain.EventUseCase {
	return &eventUseCase{
		repo: repo,
	}
}

func (uc *eventUseCase) ProcessEvent(rawEvent map[string]interface{}) (*domain.Event, error) {
	event := &domain.Event{
		Event:       rawEvent["ev"].(string),
		EventType:   rawEvent["et"].(string),
		AppID:       rawEvent["id"].(string),
		UserID:      rawEvent["uid"].(string),
		MessageID:   rawEvent["mid"].(string),
		PageTitle:   rawEvent["t"].(string),
		PageURL:     rawEvent["p"].(string),
		BrowserLang: rawEvent["l"].(string),
		ScreenSize:  rawEvent["sc"].(string),
		Attributes:  make(map[string]domain.Attribute),
		Traits:      make(map[string]domain.Attribute),
	}

	// Process attributes
	for k := range rawEvent {
		if strings.HasPrefix(k, "atrk") {
			num := k[4:]
			key := rawEvent[fmt.Sprintf("atrk%s", num)].(string)
			event.Attributes[key] = domain.Attribute{
				Value: rawEvent[fmt.Sprintf("atrv%s", num)].(string),
				Type:  rawEvent[fmt.Sprintf("atrt%s", num)].(string),
			}
		}
	}

	// Process traits
	for k := range rawEvent {
		if strings.HasPrefix(k, "uatrk") {
			num := k[5:]
			key := rawEvent[fmt.Sprintf("uatrk%s", num)].(string)
			event.Traits[key] = domain.Attribute{
				Value: rawEvent[fmt.Sprintf("uatrv%s", num)].(string),
				Type:  rawEvent[fmt.Sprintf("uatrt%s", num)].(string),
			}
		}
	}

	return event, nil
}

func (uc *eventUseCase) SendEvent(event *domain.Event) error {
	log.Println("Event send ", event)
	return uc.repo.SendToWebhook(event)
}
