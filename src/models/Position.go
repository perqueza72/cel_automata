package models

import . "own_interfaces"

type Position1D struct {
	position int
}

func NewPosition1D(position int) *Position1D {
	return &Position1D{
		position: position,
	}
}

func (p Position1D) Translate(q *IPosition) IPosition {

	new_position := NewPosition1D(p.position + (*q).GetPosition().(int))
	return new_position
}

func (p Position1D) GetPosition() interface{} {
	return p.position
}
