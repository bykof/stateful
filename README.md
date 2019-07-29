<h1 align="center">Welcome to stateful 👋</h1>
<p>
  <a href="https://travis-ci.org/bykof/stateful">
    <img alt="Travis CI" src="https://travis-ci.org/bykof/stateful.svg?branch=master" target="_blank" />
  </a>
  <a href="https://opensource.org/licenses/MIT">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" target="_blank" />
  </a>
</p>
<p>
  <a href="https://twitter.com/michaelbykovski">
    <img alt="Twitter: michaelbykovski" src="https://img.shields.io/twitter/follow/michaelbykovski.svg?style=social" target="_blank" />
  </a>
</p>

<p align="center">
  <img align="center" width="160px" src="https://github.com/bykof/stateful/raw/master/docs/resources/gopher.png">
</p>


> Create easy state machines with your existing code

# Table of Contents
1. [Documentation](#documentation)
2. [Usage](#usage)
3. [Draw graph](#draw-graph)
4. [Wildcards](#wildcards)
5. [Examples](#examples)
6. [Credits](#credits)

## Documentation

You can find the documentation here: [https://godoc.org/github.com/bykof/stateful](https://godoc.org/github.com/bykof/stateful)

## Usage 

It is very easy to use stateful.
Just create a struct and implement the `stateful` interface
```go
import "github.com/bykof/stateful"

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

// GetState and SetState implement interface stateful 
func (mm MyMachine) GetState() stateful.State {
    return mm.state
}

// GetState and SetState implement interface stateful
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
// Declare a transition of you machine and return the new state of the machine. 
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

// The state machine will check, if you transfer to a proper and defined state in the machine. See below. 
func (mm *MyMachine) FromAToNotExistingC(_ stateful.TransitionArgs) (stateful.State, error) {
	return stateful.DefaultState("C")
}
```

And now initialize the machine:
```go
myMachine := NewMyMachine()
stateMachine := &stateful.StateMachine{
    StatefulObject: &myMachine,
}

stateMachine.AddTransition(
    // The transition function 
    myMachine.FromAToB,
    // SourceStates
    stateful.States{A},
    // DestinationStates
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
    // The transition function
    myMachine.FromAToB, 
    // The transition params which will be passed to the transition function
    stateful.TransitionArgs(AmountParams{Amount: 1}),
)

_ = stateMachine.Run(
    myMachine.FromBToA, 
    stateful.TransitionArgs(AmountParams{Amount: 1}),
)

err := stateMachine.Run(
   myMachine.FromBToA, 
   stateful.TransitionArgs(AmountParams{Amount: 1}),
)

// We cannot run the transition "FromBToA" from current state "A"... 
if err != nil {
    log.Fatal(err) // will print: you cannot run FromAToB from state A
}

// We cannot transfer the machine with current transition to returned state "C"
err = stateMachine.Run(
    myMachine.FromAToNotExistingC, 
    stateful.TransitionArgs(nil),
)

if err != nil {
    log.Fatal(err) // will print: you cannot transfer to state C 
}
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

## Wildcards

You can also address wildcards as SourceStates or DestinationStates

```
stateMachine.AddTransition(
    myMachine.FromBToAllStates,
    stateful.States{B},
    stateful.States{stateful.AllStates},
)
```

This will give you the opportunity to jump e.g. B to AllStates.

*Keep in mind that `AllStates` creates a lot of complexity and maybe a missbehavior. 
So use it only if you are knowing what you are doing*
 
## Examples

Have a look at the examples: [examples](https://github.com/bykof/stateful/tree/master/examples) 


## Credits

Thank you [calhoun](https://www.calhoun.io/) for the sweet gopher image! 

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