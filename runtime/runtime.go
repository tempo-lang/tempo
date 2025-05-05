package runtime

type Env struct {
	trans Transport
}

type Transport interface {
	Send(receiver string, value any)
	Recv(sender string) *Async
}

func New(trans Transport) *Env {
	return &Env{
		trans: trans,
	}
}

func (e *Env) Send(receiver string, value any) {
	e.trans.Send(receiver, value)
}

func (e *Env) Recv(sender string) *Async {
	return e.trans.Recv(sender)
}

// Async

type Async struct {
	value   any
	valChan <-chan (any)
}

func (a *Async) Get() any {
	if a.value != nil {
		return a.value
	}
	a.value = <-a.valChan
	return a.value
}

func FixedAsync(value any) *Async {
	return &Async{
		value:   value,
		valChan: nil,
	}
}

func NewAsync(valChan <-chan (any)) *Async {
	return &Async{
		value:   nil,
		valChan: valChan,
	}
}
