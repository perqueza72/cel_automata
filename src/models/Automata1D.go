package models

import (
	. "own_interfaces"
)

type Automaton1D struct {
	Board   []ICell
	Pattern IPattern
}

type Automaton2D struct {
	Board   [][]ICell
	Pattern IPattern
}
