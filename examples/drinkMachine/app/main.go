package main

import (
	"github.com/bykof/stateful"
	"github.com/bykof/stateful/examples/drinkMachine"
	"github.com/bykof/stateful/statefulGraph"
)

func main() {
	drinkMachineObject := drinkMachine.NewMachine()
	stateMachine := &stateful.StateMachine{
		StatefulObject: &drinkMachineObject,
	}
	stateMachine.AddTransition(
		drinkMachineObject.InsertCoin,
		stateful.States{
			drinkMachine.Ready,
		},
		stateful.States{
			drinkMachine.CollectedEnoughMoney,
		},
	)

	stateMachine.AddTransition(
		drinkMachineObject.Cancel,
		stateful.States{
			stateful.AllStates,
		},
		stateful.States{
			drinkMachine.Cancelled,
		},
	)

	stateMachine.AddTransition(
		drinkMachineObject.DropDrink,
		stateful.States{
			drinkMachine.CollectedEnoughMoney,
		},
		stateful.States{
			drinkMachine.DroppedDrink,
		},
	)

	stateMachine.AddTransition(
		drinkMachineObject.DropChange,
		stateful.States{
			drinkMachine.Cancelled,
			drinkMachine.DroppedDrink,
		},
		stateful.States{
			drinkMachine.Ready,
		},
	)

	stateMachine.AddTransition(
		drinkMachineObject.GoToMaintenance,
		stateful.States{
			drinkMachine.Ready,
		},
		stateful.States{
			drinkMachine.Maintenance,
		},
	)

	stateMachine.AddTransition(
		drinkMachineObject.GoToAny,
		stateful.States{
			drinkMachine.Maintenance,
		},
		stateful.States{
			stateful.AllStates,
		},
	)
	stateMachineGraph := statefulGraph.StateMachineGraph{StateMachine: *stateMachine}
	_ = stateMachineGraph.DrawGraph()
}
