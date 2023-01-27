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
		timelaps.Automatas[len(timelaps.Automatas)-1] = (*automata).Copy()

		new_automata := (*automata).Transition()
		timelaps.Automatas = append(timelaps.Automatas, new_automata)
		return new_automata
	}

	*automata = *timelaps.Automatas[(*automata).GetId()+1]

	return automata
}

func (timeLaps *TimeLaps) Previous() *IAutomataCellular {
	automata := timeLaps.actual_automata

	if (*automata).GetId() == 0 {
		return timeLaps.actual_automata
	}
	*automata = *timeLaps.Automatas[(*automata).GetId()-1]

	return automata
}

func (timeLaps TimeLaps) Get(id uint) *IAutomataCellular {

	if int(id) < len(timeLaps.Automatas) {
		*timeLaps.actual_automata = *timeLaps.Automatas[id]
		return timeLaps.actual_automata
	}

	return timeLaps.actual_automata
}
