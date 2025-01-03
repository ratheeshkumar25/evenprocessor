package domain

type EventRepository interface {
	SendToWebhook(*Event) error
}

type EventUseCase interface {
	ProcessEvent(rawEvent map[string]interface{}) (*Event, error)
	SendEvent(event *Event) error
}

type EventWorker interface {
	ProcessEvent(event *Event)
	Start()
	Stop()
}
