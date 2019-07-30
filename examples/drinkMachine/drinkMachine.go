package drinkMachine

import (
	"github.com/bykof/stateful"
	"github.com/pkg/errors"
)

const (
	PricePerDrink        = 10
	Ready                = stateful.DefaultState("Ready")
	CollectedEnoughMoney = stateful.DefaultState("CollectedEnoughMoney")
	DroppedDrink         = stateful.DefaultState("DroppedDrink")
	Cancelled            = stateful.DefaultState("Cancelled")
	Maintenance          = stateful.DefaultState("Maintenance")
)

type (
	DrinkMachine struct {
		state          stateful.State
		currentAmount  int
		collectedCoins int
	}

	InsertCoinParam struct {
		Amount int
	}
)

func (cp *InsertCoinParam) GetData() interface{} {
	return cp
}

func NewMachine() DrinkMachine {
	return DrinkMachine{state: stateful.DefaultState(Ready)}
}

func (m DrinkMachine) State() stateful.State {
	return m.state
}

func (m *DrinkMachine) SetState(state stateful.State) error {
	m.state = state
	return nil
}

func (m DrinkMachine) GetCurrentAmount() int {
	return m.currentAmount
}

func (m DrinkMachine) GetCollectedCoins() int {
	return m.collectedCoins
}

func (m *DrinkMachine) InsertCoin(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	coinParams, ok := transitionArguments.(InsertCoinParam)
	if !ok {
		return nil, errors.New("cannot parse coinparams")
	}

	m.currentAmount += coinParams.Amount
	if m.currentAmount >= PricePerDrink {
		return CollectedEnoughMoney, nil
	}
	return m.state, nil
}

func (m *DrinkMachine) Cancel(_ stateful.TransitionArguments) (stateful.State, error) {
	return Cancelled, nil
}

func (m *DrinkMachine) GoToMaintenance(_ stateful.TransitionArguments) (stateful.State, error) {
	return Maintenance, nil
}

func (m *DrinkMachine) GoToAny(_ stateful.TransitionArguments) (stateful.State, error) {
	return Ready, nil
}

func (m *DrinkMachine) DropChange(_ stateful.TransitionArguments) (stateful.State, error) {
	m.currentAmount = 0
	return Ready, nil
}

func (m *DrinkMachine) DropDrink(_ stateful.TransitionArguments) (stateful.State, error) {
	m.collectedCoins += PricePerDrink
	m.currentAmount -= PricePerDrink
	return DroppedDrink, nil
}

func (m *DrinkMachine) NotAvailable(_ stateful.TransitionArguments) (stateful.State, error) {
	return DroppedDrink, nil
}
