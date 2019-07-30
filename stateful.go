package stateful

type (
	Stateful interface {
		State() State
		SetState(state State) error
	}
)
