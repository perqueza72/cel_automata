package logic

import . "own_interfaces"

type Position1D struct {
	Position int
}

func NewPosition1D(position int) *Position1D {
	return &Position1D{
		Position: position,
	}
}

func (p Position1D) Translate(q *IPosition) IPosition {

	new_position := NewPosition1D(p.Position + (*q).GetPosition().(int))
	return new_position
}

func (p Position1D) GetPosition() interface{} {
	return p.Position
}

////////////////////////////////////////////////////////////////////////
// Position 2D
////////////////////////////////////////////////////////////////////////

type Position2D struct {
	Row, Col int
}

func NewPosition2D(row, col int) *Position2D {
	return &Position2D{
		Row: row,
		Col: col,
	}
}

func (p Position2D) GetPosition() interface{} {
	return [2]int{p.Row, p.Col}
}

func (p Position2D) Translate(q *IPosition) IPosition {
	position := (*q).GetPosition().([2]int)
	neo_row := p.Row + position[0]
	neo_col := p.Col + position[1]

	return NewPosition2D(neo_row, neo_col)
}
