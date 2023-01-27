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
