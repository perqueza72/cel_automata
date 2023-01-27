package logic

import . "own_interfaces"

type Cell struct {
	state    bool
	position IPosition
}

func NewCell(state bool, p IPosition) *Cell {
	return &Cell{
		state:    state,
		position: p,
	}
}

func (cell *Cell) changeState() {
	cell.state = !cell.state
}

func (cell *Cell) GetPosition() *IPosition {
	return &cell.position
}

func (cell *Cell) GetState() bool {
	return cell.state
}

func (cell *Cell) SetState(state bool) {
	cell.state = state
}
