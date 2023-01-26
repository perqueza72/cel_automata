package models

import . "own_interfaces"

type Cell1D struct {
	state    bool
	position IPosition
}

func NewCell1D(state bool, p IPosition) *Cell1D {
	return &Cell1D{
		state:    state,
		position: p,
	}
}

func (cell *Cell1D) changeState() {
	cell.state = !cell.state
}

func (cell *Cell1D) GetPosition() *IPosition {
	return &cell.position
}

func (cell *Cell1D) GetState() bool {
	return cell.state
}

func (cell *Cell1D) SetState(state bool) {
	cell.state = state
}
