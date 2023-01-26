package automatas

import (
	models "automata_models"
	"fmt"
	. "own_interfaces"
)

type Automata1D struct {
	board   []ICell
	pattern *IPattern
	id      uint
}

func NewAutomata(board []ICell, pattern *IPattern, id uint) *Automata1D {
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

func Copy(a *Automata1D) *Automata1D {
	neo_board := make([]ICell, 0)

	for i := 0; i < len(a.board); i++ {
		neo_board = append(neo_board, models.NewCell1D(a.board[i].GetState(), *a.board[i].GetPosition()))
	}
	return NewAutomata(neo_board, a.pattern, a.id)
}

func (automata *Automata1D) Transition() *IAutomataCellular {
	prev_automata := Copy(automata)

	for _, cell := range automata.board {

		state := (*automata.pattern).Check(prev_automata, &cell)
		cell.SetState(state)
	}

	r := IAutomataCellular(automata)
	return &r
}

func (automata Automata1D) GetBoard() interface{} {

	states := make([]bool, 0)

	for _, v := range automata.board {
		fmt.Printf("%v", v.GetState())
		states = append(states, v.GetState())
	}

	return states
}
