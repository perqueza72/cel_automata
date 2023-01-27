package mappers

import (
	automatas "automata_automatas"
	models "models_automata"
	. "own_interfaces"
)

func ModelAutomata1D2Logic(automata *models.Automaton1D) IAutomataCellular {
	return automatas.NewAutomata1D(automata.Board, &automata.Pattern, 0)
}

func ModelAutomata2D2Logic(automata *models.Automaton2D) IAutomataCellular {
	return automatas.NewAutomata2D(automata.Board, &automata.Pattern, 0)
}
