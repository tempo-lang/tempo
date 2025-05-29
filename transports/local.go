package transports

import (
	"sync"
	"tempo/runtime"
)

type localChan struct {
	lock sync.Mutex
	send []any
	recv []chan any
}

func newLocalChan() *localChan {
	return &localChan{
		send: []any{},
		recv: []chan any{},
	}
}

func (l *localChan) Send(value any) {
	l.lock.Lock()
	if len(l.recv) > 0 {
		receiver := l.recv[0]
		l.recv = l.recv[1:]
		l.lock.Unlock()
		receiver <- value
	} else {
		l.send = append(l.send, value)
		l.lock.Unlock()
	}
}

func (l *localChan) Recv() *runtime.Async[any] {
	l.lock.Lock()
	defer l.lock.Unlock()

	if len(l.send) > 0 {
		value := l.send[0]
		l.send = l.send[1:]
		return runtime.FixedAsync(value)
	} else {
		value := make(chan any, 1)
		l.recv = append(l.recv, value)
		return runtime.NewAsync(func() any {
			return <-value
		})
	}
}

type LocalQueue struct {
	channels map[string]*localChan
	chanLock sync.Mutex
}

func NewLocal() *LocalQueue {
	return &LocalQueue{
		channels: map[string]*localChan{},
	}
}

func (q *LocalQueue) Role(role string) runtime.Transport {
	return &localTransport{
		role:  role,
		queue: q,
	}
}

func (q *LocalQueue) get(from, to string) *localChan {
	key := from + "." + to

	q.chanLock.Lock()
	if _, keyExists := q.channels[key]; !keyExists {
		q.channels[key] = newLocalChan()
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
func (l *localTransport) Recv(role string, value any) *runtime.Async[any] {
	channel := l.queue.get(role, l.role)
	return channel.Recv()
}

// Send implements runtime.Transport.
func (l *localTransport) Send(value any, roles ...string) {
	for _, receiver := range roles {
		channel := l.queue.get(l.role, receiver)
		channel.Send(value)
	}
}
