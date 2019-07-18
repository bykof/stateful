package stateful

type (
	TransitionRule struct {
		SourceStates      States
		Transition        Transition
		DestinationStates States
	}
	TransitionRules []TransitionRule
)

func (tr TransitionRule) IsAllowedToRun(state State) bool {
	return tr.SourceStates.Contains(state) || tr.SourceStates.HasWildCard()
}

func (tr TransitionRule) IsAllowedToTransfer(state State) bool {
	return tr.DestinationStates.Contains(state) || tr.DestinationStates.HasWildCard()
}

func (trs TransitionRules) Find(transition Transition) *TransitionRule {
	for _, transitionRule := range trs {
		if transitionRule.Transition.GetID() == transition.GetID() {
			return &transitionRule
		}
	}
	return nil
}