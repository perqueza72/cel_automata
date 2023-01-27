package mappers

import (
	automatas "automata_automatas"
	logic "automata_logic"
	models "models_automata"
	. "own_interfaces"
)

func ModelAutomata1D2Logic(automata *models.Automaton1D) IAutomataCellular {

	board := []ICell{}
	for i, v := range automata.Board {
		position := logic.NewPosition1D(i)
		cell := logic.NewCell(v, position)
		board = append(board, cell)
	}

	cells := []*ICell{}
	for _, v := range automata.Pattern.Cells {
		cell := ICell(logic.NewCell(v.State, v.Position))
		cells = append(cells, &cell)
	}

	pattern := IPattern(logic.NewPattern(cells))

	return automatas.NewAutomata1D(board, &pattern, 0)
}

func ModelAutomata2D2Logic(automata *models.Automaton2D) IAutomataCellular {

	board := [][]ICell{}
	for i, u := range automata.Board {
		row := []ICell{}
		for j, v := range u {
			position := logic.NewPosition2D(i, j)
			cell := logic.NewCell(v, position)
			row = append(row, cell)
		}
		board = append(board, row)
	}

	cells := []*ICell{}
	for _, v := range automata.Pattern.Cells {
		cell := ICell(logic.NewCell(v.State, v.Position))
		cells = append(cells, &cell)
	}

	pattern := IPattern(logic.NewPattern(cells))

	return automatas.NewAutomata2D(board, &pattern, 0)
}
