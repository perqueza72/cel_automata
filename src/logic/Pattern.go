package logic

import . "own_interfaces"

type Pattern struct {
	cells []*ICell
}

func NewPattern(cells []*ICell) *Pattern {
	return &Pattern{
		cells: cells,
	}
}

func (pattern *Pattern) Check(automata IAutomataCellular, cell *ICell) bool {
	for _, pattern_cell := range pattern.cells {
		position := (*(*pattern_cell).GetPosition()).Translate((*cell).GetPosition())

		current_cell, exist := automata.GetCell(position)

		if !exist {
			continue
		}

		current_state := (*current_cell).GetState()

		if (*pattern_cell).GetState() != current_state {
			return false
		}
	}

	return true
}
