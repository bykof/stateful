package stateful

import (
	"reflect"
	"runtime"
	"strings"
)

type (
	// TransitionArguments represents the arguments
	TransitionArguments interface{}

	// Transition represents the transition function which will be executed if the order is in the proper state
	// and there is a valid transitionRule in the state machine
	Transition func(transitionArguments TransitionArguments) (State, error)

	// Transitions are a slice of Transition
	Transitions []Transition
)

func (t Transition) GetName() string {
	name := runtime.FuncForPC(reflect.ValueOf(t).Pointer()).Name()
	splittedName := strings.Split(name, ".")
	splittedActualName := strings.Split(splittedName[len(splittedName)-1], "-")
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
