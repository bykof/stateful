package stateful

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TransitionTestStruct struct{}

func (tt TransitionTestStruct) abc(params TransitionArguments) (State, error) {
	return nil, nil
}

func transitiontestA(_ TransitionArguments) (State, error) {
	return nil, nil
}
func transitiontestB(_ TransitionArguments) (State, error) {
	return nil, nil
}
func transitiontestC(_ TransitionArguments) (State, error) {
	return nil, nil
}

func TestTransition_GetName(t *testing.T) {
	transitionTestStruct := TransitionTestStruct{}
	assert.Equal(t, "transitiontestA", Transition(transitiontestA).GetName())
	assert.Equal(t, "transitiontestB", Transition(transitiontestB).GetName())
	assert.Equal(t, "transitiontestC", Transition(transitiontestC).GetName())
	assert.Equal(t, "abc", Transition(transitionTestStruct.abc).GetName())
}

func TestTransitions_Contains(t *testing.T) {
	transitions := Transitions{
		Transition(transitiontestA),
		Transition(transitiontestB),
	}

	assert.True(t, transitions.Contains(Transition(transitiontestA)))
	assert.True(t, transitions.Contains(Transition(transitiontestB)))
	assert.False(t, transitions.Contains(Transition(transitiontestC)))
}
