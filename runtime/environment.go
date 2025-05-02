package runtime

type Env struct {
}

func (e Env) Send(token int, value any) {}

func (e Env) Recv(token int) any {
	return nil
}

func (e Env) EnteredRoutine() Env {
	return e
}
