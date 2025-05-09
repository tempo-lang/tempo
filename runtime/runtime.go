package runtime

type Env struct {
	trans   Transport
	roleSub map[string]string
}

type Transport interface {
	Send(value any, roles ...string)
	Recv(role string) *Async
}

func New(trans Transport) *Env {
	return &Env{
		trans: trans,
	}
}

func (e *Env) Send(value any, roles ...string) {
	subRoles := make([]string, len(roles))
	for i, role := range roles {
		if sub, ok := e.roleSub[role]; ok {
			subRoles[i] = sub
		}
	}
	e.trans.Send(value, subRoles...)
}

func (e *Env) Recv(role string) *Async {
	if sub, ok := e.roleSub[role]; ok {
		role = sub
	}
	return e.trans.Recv(role)
}

// Role maps a static role name to the name substituted in the invocation of the current function.
func (e *Env) Role(name string) string {
	return e.roleSub[name]
}

// Substitute will return a copy of the environment with a new role substitution map.
func (e *Env) Substitute(roles ...string) *Env {
	newSub := map[string]string{}
	for i := 0; i < len(roles)/2; i += 2 {
		old := roles[i]
		new := roles[i+1]
		newSub[old] = new
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
