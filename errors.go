package stateful

import (
	"fmt"
)

type (
	TransitionRuleNotFoundError struct {
		transition Transition
	}

	CannotRunFromStateError struct {
		stateMachine   StateMachine
		transitionRule TransitionRule
	}

	CannotTransferToStateError struct {
		state State
	}
)

func NewTransitionRuleNotFoundError(transition Transition) *TransitionRuleNotFoundError {
	return &TransitionRuleNotFoundError{
		transition: transition,
	}
}

func NewCannotRunFromStateError(stateMachine StateMachine, transitionRule TransitionRule) *CannotRunFromStateError {
	return &CannotRunFromStateError{
		stateMachine:   stateMachine,
		transitionRule: transitionRule,
	}
}

func NewCannotTransferToStateError(state State) *CannotTransferToStateError {
	return &CannotTransferToStateError{
		state: state,
	}
}

func (trnfe TransitionRuleNotFoundError) Error() string {
	return fmt.Sprintf(
		"no transitionRule found for given transition %s",
		trnfe.transition.GetName(),
	)
}

func (crfse CannotRunFromStateError) Error() string {
	return fmt.Sprintf(
		"you cannot run %s from state %s",
		crfse.transitionRule.Transition.GetName(),
		crfse.stateMachine.StatefulObject.GetState(),
	)
}

func (cttse CannotTransferToStateError) Error() string {
	return fmt.Sprintf(
		"you cannot transfer to state %s",
		cttse.state,
	)
}
