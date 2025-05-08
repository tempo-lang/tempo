package types

import "fmt"

type Async struct {
	inner Value
}

func (a *Async) IsSendable() bool {
	return false
}

func (a *Async) IsValue() {}

func (a *Async) ToString() string {
	return fmt.Sprintf("async %s", a.inner.ToString())
}

func (a *Async) Inner() Value {
	return a.inner
}

func NewAsync(inner Value) Value {
	if _, ok := inner.(*Async); ok {
		panic("nested async type")
	}

	return &Async{inner: inner}
}
