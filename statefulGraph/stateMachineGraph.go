package statefulGraph

import (
	"fmt"
	"github.com/awalterschulze/gographviz"
	"github.com/bykof/stateful"
)

type StateMachineGraph struct {
	StateMachine stateful.StateMachine
}

func (smg StateMachineGraph) DrawStates(graph *gographviz.Graph) error {
	for _, state := range smg.StateMachine.GetAllStates() {
		err := graph.AddNode(state.GetID(), state.GetID(), map[string]string{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (smg StateMachineGraph) DrawEdges(graph *gographviz.Graph) error {
	allStates := smg.StateMachine.GetAllStates()
	for _, transitionRule := range smg.StateMachine.GetTransitionRules() {
		sourceStates := transitionRule.SourceStates
		if sourceStates.HasWildCard() {
			sourceStates = allStates
		}

		destinationStates := transitionRule.DestinationStates
		if destinationStates.HasWildCard() {
			destinationStates = allStates
		}

		for _, sourceState := range sourceStates {
			for _, destinationState := range destinationStates {
				err := graph.AddEdge(
					sourceState.GetID(),
					destinationState.GetID(),
					true,
					map[string]string{
						"label": fmt.Sprint("\"", transitionRule.Transition.GetName(), "\""),
					},
				)

				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (smg StateMachineGraph) DrawGraph() error {
	var err error
	graph, err := intializeGraphWithDir()
	if err != nil {
		return err
	}

	err = smg.drawGraph(graph)
	if err != nil {
		return err
	}

	fmt.Println(graph.String())

	return nil
}

func (smg StateMachineGraph) DrawGraphWithName(name string) error {
	var err error
	graph, err := intializeGraphWithDir()
	if err != nil {
		return err
	}

	err = graph.SetName(name)
	if err != nil {
		return err
	}

	err = smg.drawGraph(graph)
	if err != nil {
		return err
	}

	fmt.Println(graph.String())

	return nil
}

func intializeGraphWithDir() (*gographviz.Graph, error) {
	var err error
	graph := gographviz.NewGraph()

	err = graph.SetDir(true)
	if err != nil {
		return nil, err
	}
	return graph, nil
}

func (smg StateMachineGraph) drawGraph(graph *gographviz.Graph) error {
	err := smg.DrawStates(graph)
	if err != nil {
		return err
	}

	err = smg.DrawEdges(graph)
	if err != nil {
		return err
	}
	return nil
}
