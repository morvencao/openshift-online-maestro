package event

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/openshift-online/maestro/pkg/api"
	"k8s.io/klog/v2"
)

// resourceHandler is a function that can handle resource/filesyncer status change events.
type resourceHandler func(obj interface{}) error

// eventClient is a client that can receive and handle resource/filesyncer status change events.
type eventClient struct {
	source  string
	handler resourceHandler
	errChan chan<- error
}

// EventBroadcaster is a component that can broadcast resource/filesyncer status change events to registered clients.
type EventBroadcaster struct {
	mu sync.RWMutex

	// registered clients.
	clients map[string]*eventClient

	// inbound messages from the clients.
	broadcast chan interface{}
}

// NewEventBroadcaster creates a new event broadcaster.
func NewEventBroadcaster() *EventBroadcaster {
	return &EventBroadcaster{
		clients:   make(map[string]*eventClient),
		broadcast: make(chan interface{}),
	}
}

// Register registers a client and return client id and error channel.
func (h *EventBroadcaster) Register(source string, handler resourceHandler) (string, <-chan error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	id := uuid.NewString()
	errChan := make(chan error)
	h.clients[id] = &eventClient{
		source:  source,
		handler: handler,
		errChan: errChan,
	}

	klog.V(4).Infof("register a broadcaster client %s (source=%s)", id, source)

	return id, errChan
}

// Unregister unregisters a client by id
func (h *EventBroadcaster) Unregister(id string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	close(h.clients[id].errChan)
	delete(h.clients, id)
}

// Broadcast broadcasts a resource/filesyncer status change event to all registered clients.
func (h *EventBroadcaster) Broadcast(obj interface{}) {
	h.broadcast <- obj
}

// Start starts the event broadcaster and waits for events to broadcast.
func (h *EventBroadcaster) Start(ctx context.Context) {
	klog.Infof("Starting event broadcaster")

	for {
		select {
		case <-ctx.Done():
			return
		case obj := <-h.broadcast:
			switch res := obj.(type) {
			case *api.Resource:
				h.mu.RLock()
				for _, client := range h.clients {
					if client.source == res.Source {
						if err := client.handler(res); err != nil {
							client.errChan <- err
						}
					}
				}
				h.mu.RUnlock()
			case *api.FileSyncer:
				h.mu.RLock()
				for _, client := range h.clients {
					if client.source == res.Source {
						if err := client.handler(res); err != nil {
							client.errChan <- err
						}
					}
				}
				h.mu.RUnlock()
			default:
				klog.Errorf("unknown event type: %T", res)
			}
		}
	}
}
