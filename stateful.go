package stateful

type (
	Stateful interface {
		GetState() State
		SetState(state State) error
	}
)
