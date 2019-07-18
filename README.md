<h1 align="center">Welcome to stateful 👋</h1>
<p>
  <a href="https://opensource.org/licenses/MIT">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" target="_blank" />
  </a>
  <a href="https://twitter.com/michaelbykovski">
    <img alt="Twitter: michaelbykovski" src="https://img.shields.io/twitter/follow/michaelbykovski.svg?style=social" target="_blank" />
  </a>
</p>

> Create easy state machines with your existing code

## Usage 

It is very easy to use stateful.
Just create a struct and implement the `stateful` interface
```go
import "github.com/bykof/stateful/src/stateful"

type (
	MyMachine struct {
        state   stateful.State
        amount  int
    }
	
	AmountParams struct {
        Amount  int
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
```

Declare some proper states:
```go
const (
	A = stateful.DefaultState("A")
	B = stateful.DefaultState("B")
)
```

Then add some transitions to the machine:
```go
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
```

And now initialize the machine:
```go
myMachine := NewMyMachine()
stateMachine := &stateful.StateMachine{
    StatefulObject: &myMachine,
}

stateMachine.AddTransition(
    myMachine.FromAToB,
    stateful.States{A},
    stateful.States{B},
)

stateMachine.AddTransition(
    myMachine.FromBToA,
    stateful.States{B},
    stateful.States{A},
)
```

Everything is done! Now run the machine:
```go
_ := stateMachine.Run(
	myMachine.FromAToB, 
	stateful.TransitionArgs(AmountParams{Amount: 1}),
)

_ = stateMachine.Run(
	myMachine.FromBToA, 
	stateful.TransitionArgs(AmountParams{Amount: 1}),
)
```

That's it!

## Draw graph

You can draw a graph of your state machine in `dot` format of graphviz.

Just pass in your created statemachine into the StateMachineGraph.

```go
import "github.com/bykof/stateful/src/statefulGraph"
stateMachineGraph := statefulGraph.StateMachineGraph{StateMachine: *stateMachine}
_ = stateMachineGraph.DrawGraph()
```

This will print following to the console:
```
digraph  {
	A->B[ label="FromAToB" ];
	B->A[ label="FromBToA" ];
	A;
	B;
	
}
```

which is actually this graph:

![MyMachine Transition Graph](https://github.com/bykof/stateful/raw/master/docs/resources/myMachine.png)

## Run tests

```sh
go test ./...
```

## Author

👤 **Michael Bykovski**

* Twitter: [@michaelbykovski](https://twitter.com/michaelbykovski)
* Github: [@bykof](https://github.com/bykof)

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2019 [Michael Bykovski](https://github.com/bykof).<br />
This project is [MIT](https://opensource.org/licenses/MIT) licensed.