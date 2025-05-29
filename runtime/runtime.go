package runtime

type Env struct {
	trans   Transport
	roleSub map[string]string
}

type Transport interface {
	Send(value any, roles ...string)
	Recv(role string, value any) *Async[any]
}

func New(trans Transport) *Env {
	return &Env{
		trans: trans,
	}
}

func Send[T any](env *Env, value T, roles ...string) {
	subRoles := make([]string, len(roles))
	for i, role := range roles {
		subRoles[i] = env.Role(role)
	}
	env.trans.Send(value, subRoles...)
}

func Recv[T any](env *Env, role string) *Async[T] {
	role = env.Role(role)
	var value T
	return DowncastAsync[T](env.trans.Recv(role, value))
}

// Role maps a static role name to the name substituted in the invocation of the current function.
func (e *Env) Role(name string) string {
	if sub, ok := e.roleSub[name]; ok {
		return sub
	} else {
		return name
	}
}

// Subst will return a copy of the environment with a new role substitution map.
func (e *Env) Subst(roles ...string) *Env {
	newSub := map[string]string{}
	for i := 0; i < len(roles); i += 2 {
		old := roles[i]
		new := roles[i+1]
		newSub[new] = e.Role(old)
	}

	return &Env{
		trans:   e.trans,
		roleSub: newSub,
	}
}

// Clone will return a copy of the environment.
func (e *Env) Clone() *Env {
	newRoleSub := map[string]string{}
	for key, value := range e.roleSub {
		newRoleSub[key] = value
	}

	return &Env{
		trans:   e.trans,
		roleSub: newRoleSub,
	}
}

// Async

type Async[T any] struct {
	value    T
	callback func() T
	called   bool
}

func GetAsync[T any](a *Async[T]) T {
	if a.called {
		return a.value
	}
	a.value = a.callback()
	a.called = true
	return a.value
}

func FixedAsync[T any](value T) *Async[T] {
	return &Async[T]{
		value:    value,
		callback: nil,
		called:   true,
	}
}

func NewAsync[T any](callback func() T) *Async[T] {
	return &Async[T]{
		callback: callback,
		called:   false,
	}
}

func DowncastAsync[T any](async *Async[any]) *Async[T] {
	if async.called {
		return FixedAsync(async.value.(T))
	} else {
		return NewAsync(func() T {
			return GetAsync(async).(T)
		})
	}
}

func DynAsync[T any](async *Async[T]) *Async[any] {
	return MapAsync(async, func(value T) any {
		return value
	})
}

func MapAsync[T any, U any](async *Async[T], mapper func(value T) U) *Async[U] {
	if async.called {
		return FixedAsync(mapper(async.value))
	} else {
		return NewAsync(func() U {
			return mapper(GetAsync(async))
		})
	}
}
