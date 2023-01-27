package models

import (
	. "automata_logic"
)

type Cell1D struct {
	State    bool
	Position Position1D
}

type Pattern1D struct {
	Cells []Cell1D
}

type Automaton1D struct {
	Board   []bool
	Pattern Pattern1D
}

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
