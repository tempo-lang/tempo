// Simulator implements a convenice environment for running processes locally.
package simulator

import (
	"sync"

	"github.com/tempo-lang/tempo/runtime"
	"github.com/tempo-lang/tempo/transports"
)

type Process struct {
	role string
	run  func(*runtime.Env) any
}

type Result struct {
	Return   any
	Sends    []transports.SendValue
	Receives []transports.RecvValue
}

func Proc(role string, run func(env *runtime.Env) any) Process {
	return Process{
		role: role,
		run:  run,
	}
}

// Run simulates the given processes locally.
func Run(processes ...Process) []Result {
	queue := transports.NewLocal()
	results := make([]Result, len(processes))
	var wg sync.WaitGroup

	for i, proc := range processes {
		trans := transports.NewRecorder(queue.Role(proc.role))

		wg.Add(1)
		go func() {
			ret := proc.run(runtime.New(trans))
			results[i] = Result{
				Return:   ret,
				Sends:    trans.SendValues(),
				Receives: trans.ReceivedValues(),
			}
			wg.Done()
		}()
	}

	wg.Wait()

	return results
}
