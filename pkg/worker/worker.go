package worker

import (
	"github.com/ratheeshkumar/event-processor/pkg/domain"
)

type eventWorker struct {
	eventChan chan *domain.Event
	useCase   domain.EventUseCase
	done      chan struct{}
}

func NewEventWorker(bufferSize int, useCase domain.EventUseCase) domain.EventWorker {
	return &eventWorker{
		eventChan: make(chan *domain.Event, bufferSize),
		useCase:   useCase,
		done:      make(chan struct{}),
	}
}

func (w *eventWorker) ProcessEvent(event *domain.Event) {
	w.eventChan <- event
}

func (w *eventWorker) Start() {
	go func() {
		for {
			select {
			case event := <-w.eventChan:
				if err := w.useCase.SendEvent(event); err != nil {
					// Log error here
				}
			case <-w.done:
				return
			}
		}
	}()
}

func (w *eventWorker) Stop() {
	close(w.done)
}
