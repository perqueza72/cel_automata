package automatas

import (
	logic "automata_logic"
	"fmt"
	. "own_interfaces"
)

type Automata1D struct {
	board   []ICell
	pattern *IPattern
	id      uint
}

func NewAutomata1D(board []ICell, pattern *IPattern, id uint) *Automata1D {
	return &Automata1D{
		board:   board,
		pattern: pattern,
		id:      id,
	}
}

func (automata *Automata1D) GetCell(pos IPosition) (*ICell, bool) {

	position := pos.GetPosition().(int)
	if 0 <= position && position < len(automata.board) {
		return &automata.board[position], true
	}
	return nil, false
}

func (automata *Automata1D) GetId() uint {
	return automata.id
}

func (a *Automata1D) Copy() *IAutomataCellular {
	neo_board := make([]ICell, 0)

	for i := 0; i < len(a.board); i++ {
		neo_board = append(neo_board, logic.NewCell(a.board[i].GetState(), *a.board[i].GetPosition()))
	}

	automata := NewAutomata1D(neo_board, a.pattern, a.id)
	r := IAutomataCellular(automata)
	return &r
}

func (automata *Automata1D) Transition() *IAutomataCellular {
	prev_automata := automata.Copy()

	for _, cell := range automata.board {

		state := (*automata.pattern).Check(*prev_automata, &cell)
		cell.SetState(state)
	}

	automata.id = automata.id + 1
	new_automata := automata.Copy()

	return new_automata
}

func (automata Automata1D) GetBoard() interface{} {

	states := make([]bool, 0)

	for _, v := range automata.board {
		fmt.Printf("%v", v.GetState())
		states = append(states, v.GetState())
	}

	return states
}
