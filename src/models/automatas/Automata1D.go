package models

import . "own_interfaces"

type Automata1D struct {
	board   []ICell
	pattern IPattern
	id      int
}

func (automata *Automata1D) GetCell(pos IPosition) ICell {

	position := pos.GetPosition().(int)
	return automata.board[position]
}

func (automata *Automata1D) Transition() error {
	prev_automata := *automata

	for _, cell := range automata.board {
		if automata.pattern.Check(&prev_automata, cell) {
			cell.Transition()
		}
	}

	automata = &prev_automata

	return nil
}
