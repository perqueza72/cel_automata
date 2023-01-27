package logic

import (
	. "own_interfaces"
)

type ITimeLaps interface {
	Next() *IAutomataCellular
	Previous() *IAutomataCellular
	Get(id uint) *IAutomataCellular
}

type TimeLaps struct {
	Automatas       []*IAutomataCellular
	actual_automata *IAutomataCellular
}

func NewTimeLaps(automata *IAutomataCellular) *TimeLaps {
	return &TimeLaps{
		Automatas:       []*IAutomataCellular{automata},
		actual_automata: automata,
	}
}

func (timelaps *TimeLaps) Next() *IAutomataCellular {
	automata := timelaps.actual_automata

	if int((*automata).GetId()+1) == len(timelaps.Automatas) {
		prev_automata := (*automata).Copy()
		timelaps.Automatas[len(timelaps.Automatas)-1] = prev_automata

		new_automata := (*automata).Transition()
		timelaps.Automatas = append(timelaps.Automatas, new_automata)
		return new_automata
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
