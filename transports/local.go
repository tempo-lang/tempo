package transports

import (
	"sync"
	"tempo/runtime"
)

type LocalQueue struct {
	channels map[string]chan any
	chanLock sync.Mutex
}

func NewLocal() *LocalQueue {
	return &LocalQueue{
		channels: map[string]chan any{},
	}
}

func (q *LocalQueue) Role(role string) runtime.Transport {
	return &localTransport{
		role:  role,
		queue: q,
	}
}

func (q *LocalQueue) get(from, to string) chan any {
	key := from + "." + to

	q.chanLock.Lock()
	if _, keyExists := q.channels[key]; !keyExists {
		q.channels[key] = make(chan any)
	}
	result := q.channels[key]
	q.chanLock.Unlock()

	return result
}

type localTransport struct {
	role  string
	queue *LocalQueue
}

// Recv implements runtime.Transport.
func (l *localTransport) Recv(role string) *runtime.Async {
	channel := l.queue.get(role, l.role)

	resultChan := make(chan any)
	go func() {
		value := <-channel
		resultChan <- value
	}()

	return runtime.NewAsync(func() any {
		return <-resultChan
	})
}

// Send implements runtime.Transport.
func (l *localTransport) Send(value any, roles ...string) {
	for _, receiver := range roles {
		channel := l.queue.get(l.role, receiver)
		channel <- value
	}
}
