package stateful

import (
	"reflect"
	"runtime"
	"strings"
)

type (
	TransitionArgs interface{}
	Transition func(params TransitionArgs) (State, error)
	Transitions []Transition
)

func (t Transition) GetName() string {
	name := runtime.FuncForPC(reflect.ValueOf(t).Pointer()).Name()
	splittedName := strings.Split(name, ".")
	splittedActualName := strings.Split(splittedName[len(splittedName) - 1], "-")
	return splittedActualName[0]
}

func (t Transition) GetID() uintptr {
	return reflect.ValueOf(t).Pointer()
}

func (ts Transitions) Contains(transition Transition) bool {
	for _, currentTransition := range ts {
		if currentTransition.GetID() == transition.GetID() {
			return true
		}
	}
	return false
}
