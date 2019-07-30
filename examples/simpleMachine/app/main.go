package main

import (
	"github.com/bykof/stateful"
	"github.com/bykof/stateful/examples/simpleMachine"
	"github.com/bykof/stateful/statefulGraph"
)

func main() {
	myMachine := simpleMachine.NewMyMachine()
	stateMachine := &stateful.StateMachine{
		StatefulObject: &myMachine,
	}

	stateMachine.AddTransition(
		myMachine.FromAToB,
		stateful.States{simpleMachine.A},
		stateful.States{simpleMachine.B},
	)

	stateMachine.AddTransition(
		myMachine.FromBToA,
		stateful.States{simpleMachine.B},
		stateful.States{simpleMachine.A},
	)

	_ = stateMachine.Run(
		myMachine.FromAToB,
		simpleMachine.AmountArguments{Amount: 1},
	)

	_ = stateMachine.Run(
		myMachine.FromBToA,
		simpleMachine.AmountArguments{Amount: 1},
	)

	stateMachineGraph := statefulGraph.StateMachineGraph{StateMachine: *stateMachine}
	_ = stateMachineGraph.DrawGraph()
}
