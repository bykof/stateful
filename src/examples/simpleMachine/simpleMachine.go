package simpleMachine

import (
	"errors"
	"stateful/src/stateful"
)

const (
	A = stateful.DefaultState("A")
	B = stateful.DefaultState("B")
)

type (
	MyMachine struct {
		state  stateful.State
		amount int
	}

	AmountParams struct {
		Amount int
	}
)

func NewMyMachine() MyMachine {
	return MyMachine{
		state:  A,
		amount: 0,
	}
}

func (mm MyMachine) GetState() stateful.State {
	return mm.state
}

func (mm *MyMachine) SetState(state stateful.State) error {
	mm.state = state
	return nil
}

func (mm *MyMachine) FromAToB(params stateful.TransitionArgs) (stateful.State, error) {
	amountParams, ok := params.(AmountParams)
	if !ok {
		return nil, errors.New("could not parse AmountParams")
	}

	mm.amount += amountParams.Amount
	return B, nil
}

func (mm *MyMachine) FromBToA(params stateful.TransitionArgs) (stateful.State, error) {
	amountParams, ok := params.(AmountParams)
	if !ok {
		return nil, errors.New("could not parse AmountParams")
	}

	mm.amount -= amountParams.Amount
	return A, nil
}
