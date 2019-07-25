package stateful

const (
	AllStates = DefaultState("*")
)

type (
	State interface {
		GetID() string
		IsWildCard() bool
	}
	States []State
	DefaultState string
)

// IsWildCard checks if the current state is a wildcard.
// So if the state stands for all possible states
func (ds DefaultState) IsWildCard() bool {
	return ds == AllStates
}

// GetID returns the string representation of the DefaultState
func (ds DefaultState) GetID() string {
	return string(ds)
}

// Contains search in States if the given state is inside the States
// It compares with GetID
func (ss States) Contains(state State) bool {
	for _, currentState := range ss {
		if currentState.GetID() == state.GetID() {
			return true
		}
	}
	return false
}

// HasWildCard checks if there is a wildcard state inside of States
func (ss States) HasWildCard() bool {
	for _, currentState := range ss {
		if currentState.IsWildCard() {
			return true
		}
	}
	return false
}
