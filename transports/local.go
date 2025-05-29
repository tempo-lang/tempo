package transports

import (
	"encoding/json"
	"fmt"
	"sync"
	"tempo/runtime"
)

type jsonData []byte

type localChan struct {
	lock sync.Mutex
	send []jsonData
	recv []chan jsonData
}

func newLocalChan() *localChan {
	return &localChan{
		send: []jsonData{},
		recv: []chan jsonData{},
	}
}

func (l *localChan) Send(value jsonData) {
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

func (l *localChan) Recv() *runtime.Async[jsonData] {
	l.lock.Lock()
	defer l.lock.Unlock()

	if len(l.send) > 0 {
		value := l.send[0]
		l.send = l.send[1:]
		return runtime.FixedAsync(value)
	} else {
		value := make(chan jsonData, 1)
		l.recv = append(l.recv, value)
		return runtime.NewAsync(func() jsonData {
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

	return runtime.MapAsync(channel.Recv(), func(data jsonData) any {
		// json.Unmarshal converts ints to float64
		if _, isInt := value.(int); isInt {
			json.Unmarshal(data, &value)
			return int(value.(float64))
		} else {
			json.Unmarshal(data, &value)
			return value
		}
	})
}

// Send implements runtime.Transport.
func (l *localTransport) Send(value any, roles ...string) {
	for _, receiver := range roles {
		channel := l.queue.get(l.role, receiver)

		result, err := json.Marshal(&value)
		if err != nil {
			panic(fmt.Sprintf("failed to json encode message to send: %#v", value))
		}

		channel.Send(result)
	}
}
