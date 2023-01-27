package auxiliar

import . "automata_automatas"
import . "own_interfaces"
import logic "auxiliar_logic"

func Copy(a *Automata1D) *Automata1D {
	neo_board := make([]ICell, 0)

	for i := 0; i < len(a.Board); i++ {
		neo_board = append(neo_board, logic.NewCell(a.Board[i].GetState(), *a.Board[i].GetPosition()))
	}
	return NewAutomata1D(neo_board, a.pattern, a.id)
}