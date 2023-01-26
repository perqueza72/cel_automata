package models

import . "own_interfaces"

type ITimeLaps interface {
	Next() IAutomataCellular
	Previous() IAutomataCellular
	Get(id uint) IAutomataCellular
}

type TimeLaps struct {
	automata []IAutomataCellular
}
