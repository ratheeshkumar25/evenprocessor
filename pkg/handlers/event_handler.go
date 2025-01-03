package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ratheeshkumar/event-processor/pkg/models"

	"github.com/ratheeshkumar/event-processor/pkg/domain"
)

type EventHandler struct {
	useCase domain.EventUseCase
	worker  domain.EventWorker
}

func NewEventHandler(useCase domain.EventUseCase, worker domain.EventWorker) *EventHandler {
	return &EventHandler{
		useCase: useCase,
		worker:  worker,
	}
}

func (h *EventHandler) HandleEvent(w http.ResponseWriter, r *http.Request) {
	var rawEvent models.IncomingEvent
	if err := json.NewDecoder(r.Body).Decode(&rawEvent); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	event, err := h.useCase.ProcessEvent(rawEvent)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to process event")
		return
	}

	// Process event asynchronously
	go h.worker.ProcessEvent(event)

	respondWithJSON(w, http.StatusAccepted, models.APIResponse{
		Success: true,
		Message: "Event accepted for processing",
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, models.APIResponse{
		Success: false,
		Error:   message,
	})
}
