package stateful

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStates_Contains(t *testing.T) {
	state1 := DefaultState("state1")
	state2 := DefaultState("state2")
	state3 := DefaultState("state3")
	states := States{
		state1,
		state2,
	}

	assert.True(t, states.Contains(state1))
	assert.True(t, states.Contains(state2))
	assert.False(t, states.Contains(state3))
}
