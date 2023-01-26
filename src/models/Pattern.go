package models

import . "own_interfaces"

type Pattern struct {
	cells []ICell
}

func (pattern *Pattern) Check(automata IAutomataCellular, cell ICell) bool {
	for _, pattern_cell := range pattern.cells {
		position := pattern_cell.GetPosition().Append(cell.GetPosition())
		current_state := automata.GetCell(position).GetState()

		if pattern_cell.GetState() != current_state {
			return false
		}
	}

	return true
}
