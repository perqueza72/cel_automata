package models

import . "own_interfaces"

type ITimeLaps interface {
	Next() *IAutomataCellular
	Previous() *IAutomataCellular
	Get(id uint) *IAutomataCellular
}

type TimeLaps struct {
	Automatas       []*IAutomataCellular
	actual_automata *IAutomataCellular
}

func (timelaps *TimeLaps) Next() *IAutomataCellular {
	automata := timelaps.actual_automata

	if automata == timelaps.Automatas[len(timelaps.Automatas)-1] {
		return (*automata).Transition()
	}

	return timelaps.Automatas[(*automata).GetId()+1]
}

func (timeLaps *TimeLaps) Previous() *IAutomataCellular {
	automata := timeLaps.actual_automata

	if (*automata).GetId() == 0 {
		return timeLaps.actual_automata
	}

	return timeLaps.Automatas[(*automata).GetId()-1]
}

func (timeLaps TimeLaps) Get(id uint) *IAutomataCellular {

	if int(id) < len(timeLaps.Automatas) {
		return timeLaps.Automatas[id]
	}

	return timeLaps.actual_automata
}
