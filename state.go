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

func (ds DefaultState) IsWildCard() bool {
	return ds == AllStates
}

func (ds DefaultState) GetID() string {
	return string(ds)
}

func (ss States) Contains(state State) bool {
	for _, currentState := range ss {
		if currentState.GetID() == state.GetID() {
			return true
		}
	}
	return false
}

func (ss States) HasWildCard() bool {
	for _, currentState := range ss {
		if currentState.IsWildCard() {
			return true
		}
	}
	return false
}