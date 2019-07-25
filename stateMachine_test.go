package stateful

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

const (
	State1 = DefaultState("State1")
	State2 = DefaultState("State2")
	State3 = DefaultState("State3")
	State4 = DefaultState("State4")
)

type (
	TestStatefulObject struct {
		state     State
		TestValue int
	}

	TestParam struct {
		Amount int
	}
)

func (tsp TestStatefulObject) GetState() State {
	return tsp.state
}

func (tsp *TestStatefulObject) SetState(state State) error {
	tsp.state = state
	return nil
}

func (tsp *TestStatefulObject) FromState1ToState2(_ TransitionArgs) (State, error) {
	return State2, nil
}

func (tsp *TestStatefulObject) FromState2ToState3(params TransitionArgs) (State, error) {
	testParam, _ := params.(TestParam)
	tsp.TestValue += testParam.Amount
	return State3, nil
}

func (tsp *TestStatefulObject) FromState3ToState1And2(_ TransitionArgs) (State, error) {
	return State1, nil
}

func (tsp *TestStatefulObject) FromState2And3To4(_ TransitionArgs) (State, error) {
	return State4, nil
}

func (tsp *TestStatefulObject) FromState4ToState1(_ TransitionArgs) (State, error) {
	return State1, nil
}

func (tsp *TestStatefulObject) ErrorBehavior(_ TransitionArgs) (State, error) {
	return nil, errors.New("there was an error")
}

func (tsp TestStatefulObject) NotExistingTransition(_ TransitionArgs) (State, error) {
	return nil, nil
}

func (tsp TestStatefulObject) FromState3ToNotExistingState(_ TransitionArgs) (State, error) {
	return DefaultState("NotExisting"), nil
}

func NewTestStatefulObject() TestStatefulObject {
	return TestStatefulObject{state: State1}
}

func NewStateMachine() StateMachine {
	testStatefulObject := NewTestStatefulObject()
	stateMachine := StateMachine{StatefulObject: &testStatefulObject}
	stateMachine.AddTransition(testStatefulObject.FromState1ToState2, States{State1}, States{State2})
	stateMachine.AddTransition(testStatefulObject.FromState2ToState3, States{State2}, States{State3})
	stateMachine.AddTransition(testStatefulObject.FromState3ToState1And2, States{State3}, States{State1, State2})
	stateMachine.AddTransition(testStatefulObject.FromState2And3To4, States{State2, State3}, States{State4})
	stateMachine.AddTransition(testStatefulObject.FromState4ToState1, States{State4}, States{State1})
	stateMachine.AddTransition(testStatefulObject.ErrorBehavior, States{AllStates}, States{State1, State2})
	stateMachine.AddTransition(testStatefulObject.FromState3ToNotExistingState, States{State3}, States{})
	return stateMachine
}

func TestStateMachine_AddTransition(t *testing.T) {
	testStatefulObject := NewTestStatefulObject()
	stateMachine := StateMachine{StatefulObject: &testStatefulObject}
	stateMachine.AddTransition(testStatefulObject.FromState1ToState2, States{State1}, States{State2})
	stateMachine.AddTransition(testStatefulObject.ErrorBehavior, States{State2, State4}, States{State2, State3})

	assert.ElementsMatch(
		t,
		States{State1},
		stateMachine.transitionRules[0].SourceStates,
	)
	assert.ElementsMatch(
		t,
		States{State2},
		stateMachine.transitionRules[0].DestinationStates,
	)

	transition := Transition(testStatefulObject.FromState1ToState2)
	assert.Equal(
		t,
		transition.GetName(),
		stateMachine.transitionRules[0].Transition.GetName(),
	)

	assert.ElementsMatch(
		t,
		States{State2, State4},
		stateMachine.transitionRules[1].SourceStates,
	)
	assert.ElementsMatch(
		t,
		States{State2, State3},
		stateMachine.transitionRules[1].DestinationStates,
	)

	transition = Transition(testStatefulObject.ErrorBehavior)
	assert.Equal(
		t,
		transition.GetName(),
		stateMachine.transitionRules[1].Transition.GetName(),
	)
}

func TestStateMachine_GetAllStates(t *testing.T) {
	testStatefulObject := NewTestStatefulObject()
	stateMachine := StateMachine{StatefulObject: &testStatefulObject}
	stateMachine.AddTransition(testStatefulObject.FromState1ToState2, States{State1}, States{State2})
	stateMachine.AddTransition(testStatefulObject.ErrorBehavior, States{State2, State4}, States{State2, State3})

	assert.ElementsMatch(
		t,
		States{
			State1,
			State2,
			State3,
			State4,
		},
		stateMachine.GetAllStates(),
	)
}

func TestStateMachine_Run(t *testing.T) {
	stateMachine := NewStateMachine()
	testStatefulObject := stateMachine.StatefulObject.(*TestStatefulObject)
	err := stateMachine.Run(
		testStatefulObject.FromState1ToState2,
		TransitionArgs(nil),
	)
	assert.NoError(t, err)
	assert.Equal(t, State2, testStatefulObject.GetState())

	err = stateMachine.Run(
		testStatefulObject.FromState2ToState3,
		TransitionArgs(TestParam{Amount: 2}),
	)

	assert.NoError(t, err)
	assert.Equal(t, State3, testStatefulObject.GetState())
	assert.Equal(t, 2, testStatefulObject.TestValue)

	err = stateMachine.Run(
		testStatefulObject.FromState4ToState1,
		TransitionArgs(nil),
	)
	assert.Error(t, err)
	assert.Equal(
		t,
		reflect.TypeOf(&CannotRunFromStateError{}),
		reflect.TypeOf(err),
	)

	err = stateMachine.Run(
		testStatefulObject.ErrorBehavior,
		TransitionArgs(nil),
	)
	assert.Error(t, err)
	assert.Equal(t, errors.New("there was an error"), err)

	err = stateMachine.Run(
		testStatefulObject.NotExistingTransition,
		nil,
	)

	assert.Error(t, err)
	assert.Equal(
		t,
		reflect.TypeOf(&TransitionRuleNotFoundError{}),
		reflect.TypeOf(err),
	)

	err = stateMachine.Run(
		testStatefulObject.FromState3ToNotExistingState,
		nil,
	)

	assert.Error(t, err)
	assert.Equal(
		t,
		reflect.TypeOf(&CannotTransferToStateError{}),
		reflect.TypeOf(err),
	)

	assert.True(t, reflect.TypeOf(&CannotTransferToStateError{}) == reflect.TypeOf(err))
}

func TestStateMachine_GetAvailableTransitions(t *testing.T) {
	stateMachine := NewStateMachine()
	availableTransitions := stateMachine.GetAvailableTransitions()
	assert.Equal(
		t,
		Transition(stateMachine.StatefulObject.(*TestStatefulObject).FromState1ToState2).GetName(),
		availableTransitions[0].GetName(),
	)
}
