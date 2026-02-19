// This package contains the runtime used by Tempo generated source code.
package runtime

import (
	"maps"
	"reflect"
	"strconv"
)

// The environment struct keeps track of the state of the choreography at a single process.
// The first argument of any generated process function will take an `Env` as its first argument.
// It is used to interact with the environment such as sending and receiving messages.
type Env struct {
	trans   Transport
	roleSub map[string]string
}

// The Transport interface specifies the methods needed in order for processes in a choreography to communicate.
//
// The [tempo/transports] package contains a set of default implementations.
type Transport interface {
	// Send is called by a process to send a value to a set of roles.
	Send(value any, roles ...string)
	// Recv is called by a process when it expects to receive a value from a role.
	//
	// The `value` argument is an empty pointer with the expected type of the message to receive.
	// It can be used when deserializing the message to the right code, for example in use with [encoding/json.Unmarshal].
	Recv(role string, value any) *Async[any]
}

// New constructs a new environment given a transport implementation.
func New(trans Transport) *Env {
	return &Env{
		trans: trans,
	}
}

// Send will use the underlying [Transport] implementation to send the value.
// The original value is returned to make it easier to use in expressions.
func Send[T any](env *Env, value T, roles ...string) *Async[T] {
	subRoles := make([]string, len(roles))
	for i, role := range roles {
		subRoles[i] = env.Role(role)
	}
	env.trans.Send(value, subRoles...)
	return FixedAsync(value)
}

// Recv will use the underlying [Transport] implementation to receive a value.
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
	maps.Copy(newRoleSub, e.roleSub)

	return &Env{
		trans:   e.trans,
		roleSub: newRoleSub,
	}
}

// MARK: Async

// An async object represents a value that is not necessarily present yet.
// The user can call [GetAsync] to wait until the underlying value is present and get it.
type Async[T any] struct {
	value    T
	callback func() T
	called   bool
}

// GetAsync will block the thread until the underlying value becomes present and return it.
func GetAsync[T any](a *Async[T]) T {
	if a.called {
		return Copy(a.value)
	}
	a.value = a.callback()
	a.called = true
	return Copy(a.value)
}

// FixedAsync wraps a value in a resolved async value.
// Calling [GetAsync] on the returned value, will immediately yield the underlying value.
func FixedAsync[T any](value T) *Async[T] {
	return &Async[T]{
		value:    value,
		callback: nil,
		called:   true,
	}
}

// NewAsync constructs an async value.
// Calling [GetAsync] on the returned value, will call the `callback` closure to obtain the underlying value.
// The `callback` is guaranteed to ever only be called once.
func NewAsync[T any](callback func() T) *Async[T] {
	return &Async[T]{
		callback: callback,
		called:   false,
	}
}

// DowncastAsync will downcast an async any type to a specific type.
// The program will panic if the any type does not match the specific type.
func DowncastAsync[T any](async *Async[any]) *Async[T] {
	if async.called {
		return FixedAsync(async.value.(T))
	} else {
		return NewAsync(func() T {
			return GetAsync(async).(T)
		})
	}
}

// DynAsync erases the type of the async.
func DynAsync[T any](async *Async[T]) *Async[any] {
	return MapAsync(async, func(value T) any {
		return value
	})
}

// MapAsync maps an async to a new value using the `mapper` function.
func MapAsync[T any, U any](async *Async[T], mapper func(value T) U) *Async[U] {
	if async.called {
		return FixedAsync(mapper(async.value))
	} else {
		return NewAsync(func() U {
			return mapper(GetAsync(async))
		})
	}
}

// MARK: Copy

// Copy makes a semantically deep copy of the value to maintain value semantics.
// Immutable values are not necessarily copied since they effectively have value semantics.
func Copy[T any](value T) T {
	v := reflect.ValueOf(value)
	t := v.Type()
	kind := t.Kind()
	switch kind {
	case reflect.Slice:
		if t.Elem().Kind() != reflect.Slice && t.Elem().Kind() != reflect.Struct {
			// return shallow copy
			sliceCopy := reflect.MakeSlice(t, v.Len(), v.Len())
			reflect.Copy(sliceCopy, v)
			return sliceCopy.Interface().(T)
		} else {
			sliceCopy := reflect.MakeSlice(t, v.Len(), v.Len())
			for i := 0; i < v.Len(); i++ {
				elemCopy := Copy(v.Index(i).Interface())
				sliceCopy.Index(i).Set(reflect.ValueOf(elemCopy))
			}
			return sliceCopy.Interface().(T)
		}
	case reflect.Struct:
		structCopy := reflect.New(t)
		for i := 0; i < v.NumField(); i++ {
			elemCopy := Copy(v.Field(i).Interface())
			structCopy.Elem().Field(i).Set(reflect.ValueOf(elemCopy))
		}
		return structCopy.Elem().Interface().(T)
	}

	// implicit copy base type
	return value
}

// MARK: Casting

func IntToString(value int) string {
	return strconv.Itoa(value)
}

func FloatToString(value float64) string {
	return strconv.FormatFloat(value, 'g', -1, 64)
}

func BoolToString(value bool) string {
	return strconv.FormatBool(value)
}
