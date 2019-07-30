package simpleMachine

import (
	"errors"
	"github.com/bykof/stateful"
)

var (
	A = stateful.DefaultState("A")
	B = stateful.DefaultState("B")
)

type (
	MyMachine struct {
		state  stateful.State
		amount int
	}

	AmountArguments struct {
		Amount int
	}
)

func NewMyMachine() MyMachine {
	return MyMachine{
		state:  A,
		amount: 0,
	}
}

func (mm MyMachine) State() stateful.State {
	return mm.state
}

func (mm *MyMachine) SetState(state stateful.State) error {
	mm.state = state
	return nil
}

func (mm *MyMachine) FromAToB(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	amountArguments, ok := transitionArguments.(AmountArguments)
	if !ok {
		return nil, errors.New("could not parse transitionarguments as amountarguments")
	}
	mm.amount += amountArguments.Amount
	return B, nil
}

func (mm *MyMachine) FromBToA(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	amountArguments, ok := transitionArguments.(AmountArguments)
	if !ok {
		return nil, errors.New("could not parse transitionarguments as amountarguments")
	}
	mm.amount -= amountArguments.Amount
	return A, nil
}
