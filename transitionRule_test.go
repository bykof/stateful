package stateful

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func transitionruletestA(_ TransitionArguments) (State, error) {
	return nil, nil
}
func transitionruletestB(_ TransitionArguments) (State, error) {
	return nil, nil
}
func transitionruletestC(_ TransitionArguments) (State, error) {
	return nil, nil
}

func TestTransitionRule_IsAllowedToRun(t *testing.T) {
	transitionRule := TransitionRule{
		SourceStates: States{
			DefaultState("transitionTest_a"),
			DefaultState("transitionTest_b"),
		},
	}
	assert.True(t, transitionRule.IsAllowedToRun(DefaultState("transitionTest_a")))
	assert.True(t, transitionRule.IsAllowedToRun(DefaultState("transitionTest_b")))
	assert.False(t, transitionRule.IsAllowedToRun(DefaultState("transitionTest_c")))
}

func TestTransitionRule_IsAllowedToTransfer(t *testing.T) {
	transitionRule := TransitionRule{
		DestinationStates: States{
			DefaultState("transitionTest_a"),
			DefaultState("transitionTest_b"),
		},
	}
	assert.True(t, transitionRule.IsAllowedToTransfer(DefaultState("transitionTest_a")))
	assert.True(t, transitionRule.IsAllowedToTransfer(DefaultState("transitionTest_b")))
	assert.False(t, transitionRule.IsAllowedToTransfer(DefaultState("transitionTest_c")))
}

func TestTransitionRules_Find(t *testing.T) {
	transitionRules := TransitionRules{
		TransitionRule{
			Transition: transitionruletestA,
		},
		TransitionRule{
			Transition: transitionruletestB,
		},
	}

	assert.Equal(
		t,
		TransitionRule{
			Transition: transitionruletestA,
		}.Transition.GetID(),
		transitionRules.Find(transitionruletestA).Transition.GetID(),
	)

	assert.Equal(
		t,
		TransitionRule{
			Transition: transitionruletestB,
		}.Transition.GetID(),
		transitionRules.Find(transitionruletestB).Transition.GetID(),
	)
	assert.Nil(
		t,
		transitionRules.Find(transitionruletestC),
	)
}
