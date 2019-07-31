package stateful

type (
	// Stateful is the core interface which should be implemented by all stateful structs.
	// If this interface is implemented by a struct it can be processed by the state machine
	Stateful interface {
		// State returns the current state of the stateful object
		State() State

		// SetState sets the state of the stateful object and returns an error if it fails
		SetState(state State) error
	}
)
