package logic

import . "own_interfaces"

func NextState(automata *IAutomataCellular) {

	(*automata).Transition()
}
