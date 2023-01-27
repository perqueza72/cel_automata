package models

import (
	. "automata_logic"
)

type Cell2D struct {
	State    bool
	Position Position2D
}

type Pattern2D struct {
	Cells []Cell2D
}

type Automaton2D struct {
	Board   [][]bool
	Pattern Pattern2D
}
