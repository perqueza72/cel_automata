package automatas

import (
	logic "automata_logic"
	. "own_interfaces"
)

type Automata2D struct {
	board   [][]ICell
	pattern *IPattern
	id      uint
}

func NewAutomata2D(board [][]ICell, pattern *IPattern, id uint) *Automata2D {
	return &Automata2D{
		board:   board,
		pattern: pattern,
		id:      id,
	}
}

func (at *Automata2D) GetId() uint {
	return at.id
}

func (at *Automata2D) in_bounds(pos *IPosition) bool {
	pos_arr := (*pos).GetPosition().([]int)
	row, col := pos_arr[0], pos_arr[1]
	if row >= 0 && row < len(at.board) && col >= 0 && col < len(at.board[0]) {
		return true
	}
	return false
}

func (at *Automata2D) current_pos(pos *IPosition) *ICell {
	pos_arr := (*pos).GetPosition().([]int)
	row, col := pos_arr[0], pos_arr[1]
	return &at.board[row][col]
}

func (at *Automata2D) GetCell(pos IPosition) (*ICell, bool) {
	if at.in_bounds(&pos) {
		return at.current_pos(&pos), true
	}
	return nil, false
}

func (at *Automata2D) Copy() *IAutomataCellular {
	neo_board := make([][]ICell, len(at.board))

	for i := 0; i < len(at.board); i++ {
		for j := 0; j < len(at.board[i]); j++ {
			state := at.board[i][j].GetState()
			position := at.board[i][j].GetPosition()
			neo_board[i] = append(neo_board[i], logic.NewCell(state, *position))
		}
	}

	n := NewAutomata2D(neo_board, at.pattern, at.id)
	automata := IAutomataCellular(n)

	return &automata
}

func (at *Automata2D) Transition() *IAutomataCellular {
	prev_automata := at.Copy()

	for i := 0; i < len(at.board); i++ {

		for j := 0; j < len(at.board[i]); j++ {
			state := (*at.pattern).Check(*prev_automata, &at.board[i][j])
			at.board[i][j].SetState(state)
		}
	}

	r := IAutomataCellular(at)
	return &r
}

func (at *Automata2D) GetBoard() interface{} {
	states := make([][]bool, len(at.board))

	for i := 0; i < len(at.board); i++ {
		for j := 0; j < len(at.board[i]); j++ {
			states[i] = append(states[i], at.board[i][j].GetState())
		}
	}

	return states
}
