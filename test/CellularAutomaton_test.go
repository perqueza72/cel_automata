package test

import (
	automatas "automata_automatas"
	models "automata_models"
	. "own_interfaces"
	"testing"
)

func initCell(pos int, state bool) ICell {

	position := models.NewPosition1D(pos)
	cell := models.NewCell1D(state, *position)
	return cell
}

func initPattern() IPattern {

	a := initCell(-1, true)
	b := initCell(1, true)

	cells := make([]*ICell, 2)
	cells[0] = &a
	cells[1] = &b
	pattern := models.NewPattern(cells)

	return pattern
}

func initAutomata() IAutomataCellular {
	states := [5]bool{true, false, true, false, false}

	cells := make([]ICell, 0)

	for i := 0; i < len(states); i++ {
		cell := initCell(i, states[i])
		cells = append(cells, cell)
	}

	pattern := initPattern()
	return automatas.NewAutomata(cells, &pattern, 0)
}

func TestAutomaton1D(t *testing.T) {

	automata := initAutomata()
	automata.Transition()
	automata.GetBoard()

	got := automata.GetBoard().([]bool)
	expected := []bool{false, true, false, false, false}

	for i := range got {
		if got[i] != expected[i] {
			t.Errorf("In index %v got %v, whanted %v", i, got[i], expected[i])
			t.Fail()
		}
	}

}
