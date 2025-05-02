package runtime

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
