package transports

import "tempo/runtime"

type SendValue struct {
	Value     any
	Receivers []string
}

type RecvValue struct {
	Value  any
	Sender string
}

type Recorder struct {
	inner    runtime.Transport
	sends    []SendValue
	receives []RecvValue
}

func NewRecorder(transport runtime.Transport) *Recorder {
	return &Recorder{
		inner:    transport,
		sends:    []SendValue{},
		receives: []RecvValue{},
	}
}

func (r *Recorder) SendValues() []SendValue {
	return r.sends
}

func (r *Recorder) ReceivedValues() []RecvValue {
	result := []RecvValue{}
	for _, val := range r.receives {
		async := val.Value.(*runtime.Async)
		result = append(result, RecvValue{
			Value:  async.Get(),
			Sender: val.Sender,
		})
	}
	return result
}

// Recv implements runtime.Transport.
func (r *Recorder) Recv(role string) *runtime.Async {
	value := r.inner.Recv(role)

	r.receives = append(r.receives, RecvValue{
		Value:  value,
		Sender: role,
	})

	return value
}

// Send implements runtime.Transport.
func (r *Recorder) Send(value any, roles ...string) {
	r.inner.Send(value, roles...)

	r.sends = append(r.sends, SendValue{
		Value:     value,
		Receivers: roles,
	})
}
