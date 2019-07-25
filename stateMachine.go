package stateful

type (
	StateMachine struct {
		StatefulObject  Stateful
		transitionRules TransitionRules
	}
)

func (sm *StateMachine) AddTransition(
	transition Transition,
	sourceStates States,
	destinationStates States,
) {
	sm.transitionRules = append(
		sm.transitionRules,
		TransitionRule{
			SourceStates:      sourceStates,
			Transition:        transition,
			DestinationStates: destinationStates,
		},
	)
}

func (sm StateMachine) GetTransitionRules() TransitionRules {
	return sm.transitionRules
}

func (sm StateMachine) GetAllStates() States {
	states := States{}
	keys := make(map[State]bool)

	for _, transitionRule := range sm.transitionRules {
		for _, state := range append(transitionRule.SourceStates, transitionRule.DestinationStates...) {
			if _, ok := keys[state]; !ok {
				keys[state] = true
				if !state.IsWildCard() {
					states = append(states, state)
				}
			}
		}
	}
	return states
}

func (sm StateMachine) Run(
	transition Transition,
	transitionArgs TransitionArgs,
) error {
	transitionRule := sm.transitionRules.Find(transition)
	if transitionRule == nil {
		return NewTransitionRuleNotFoundError(transition)
	}

	if !transitionRule.IsAllowedToRun(sm.StatefulObject.GetState()) {
		return NewCannotRunFromStateError(sm, *transitionRule)
	}

	newState, err := transition(transitionArgs)
	if err != nil {
		return err
	}

	if !transitionRule.IsAllowedToTransfer(newState) {
		return NewCannotTransferToStateError(newState)
	}

	err = sm.StatefulObject.SetState(newState)
	if err != nil {
		return err
	}
	return nil
}

func (sm StateMachine) GetAvailableTransitions() Transitions {
	transitions := Transitions{}
	for _, transitionRule := range sm.transitionRules {
		if transitionRule.IsAllowedToRun(sm.StatefulObject.GetState()) {
			transitions = append(transitions, transitionRule.Transition)
		}
	}
	return transitions
}
