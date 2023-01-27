package logic

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

////////////////////////////////////////////////////////////////////////
// Position 2D
////////////////////////////////////////////////////////////////////////

type Position2D struct {
	row, col int
}

func NewPosition2D(row, col int) *Position2D {
	return &Position2D{
		row: row,
		col: col,
	}
}

func (p Position2D) GetPosition() interface{} {
	return [2]int{p.row, p.col}
}

func (p Position2D) Translate(q *IPosition) IPosition {
	position := (*q).GetPosition().([2]int)
	neo_row := p.row + position[0]
	neo_col := p.col + position[1]

	return NewPosition2D(neo_row, neo_col)
}
