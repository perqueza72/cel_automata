package test

import (
	automatas "automata_automatas"
	logic "automata_logic"
	. "own_interfaces"
	"testing"
)

func initCell(pos int, state bool) ICell {

	position := logic.NewPosition1D(pos)
	cell := logic.NewCell(state, *position)
	return cell
}

func initPattern() IPattern {

	a := initCell(-1, true)
	b := initCell(1, true)

	cells := make([]*ICell, 2)
	cells[0] = &a
	cells[1] = &b
	pattern := logic.NewPattern(cells)

	return pattern
}

func initAutomata(states []bool) IAutomataCellular {

	cells := make([]ICell, 0)

	for i := 0; i < len(states); i++ {
		cell := initCell(i, states[i])
		cells = append(cells, cell)
	}

	pattern := initPattern()
	return automatas.NewAutomata1D(cells, &pattern, 0)
}

func InitialStateAutomaton1D() IAutomataCellular {
	return initAutomata([]bool{true, false, true, false, false})
}

func TestAutomaton1DDummie(t *testing.T) {
	automata := initAutomata([]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false})
	timeLaps := logic.NewTimeLaps(&automata)
	new_automata := timeLaps.Next()

	got := (*new_automata).GetBoard().([]bool)
	expect := []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}

	AutomatonCorrectnessChecker(t, got, expect)
}

func TestAutomaton1D(t *testing.T) {

	automata := InitialStateAutomaton1D()
	timeLaps := logic.NewTimeLaps(&automata)
	new_automata := timeLaps.Next()

	got := (*new_automata).GetBoard().([]bool)
	expected := []bool{false, true, false, false, false}

	AutomatonCorrectnessChecker(t, got, expected)
}

// func TestDestructorDeDavid1D(t *testing.T){
// 	automata := initAutomata()
// 	new_automata := aut
// }

func AutomatonCorrectnessChecker(t *testing.T, got []bool, expected []bool) {
	for i := range got {
		if got[i] != expected[i] {
			t.Errorf("In index %v got %v, whanted %v\n", i, got[i], expected[i])
			t.Fail()
		}
	}
}

func TestAutomatonTimeLapsNext(t *testing.T) {
	automata := InitialStateAutomaton1D()

	timeLaps := logic.NewTimeLaps(&automata)
	got := (*timeLaps.Next()).GetBoard().([]bool)
	expected := []bool{false, true, false, false, false}

	AutomatonCorrectnessChecker(t, got, expected)

	got = (*timeLaps.Next()).GetBoard().([]bool)
	expected = []bool{true, false, false, false, false}
	AutomatonCorrectnessChecker(t, got, expected)

	got = (*timeLaps.Next()).GetBoard().([]bool)
	expected = []bool{false, false, false, false, false}
	AutomatonCorrectnessChecker(t, got, expected)

}

func TestAutomatonTimeLapsGet(t *testing.T) {
	automata := InitialStateAutomaton1D()

	timeLaps := logic.NewTimeLaps(&automata)

	for i := 0; i < 100; i++ {
		timeLaps.Next()
	}

	got := (*timeLaps.Get(0)).GetBoard().([]bool)
	expected := []bool{true, false, true, false, false}
	AutomatonCorrectnessChecker(t, got, expected)

	got = (*timeLaps.Get(1)).GetBoard().([]bool)
	expected = []bool{false, true, false, false, false}
	AutomatonCorrectnessChecker(t, got, expected)

	got = (*timeLaps.Get(2)).GetBoard().([]bool)
	expected = []bool{true, false, false, false, false}
	AutomatonCorrectnessChecker(t, got, expected)

	got = (*timeLaps.Get(3)).GetBoard().([]bool)
	expected = []bool{false, false, false, false, false}
	AutomatonCorrectnessChecker(t, got, expected)

	got = (*timeLaps.Get(70)).GetBoard().([]bool)
	expected = []bool{false, false, false, false, false}
	AutomatonCorrectnessChecker(t, got, expected)
}

func TestAutomatonTimeLapsPreview(t *testing.T) {
	t.Log("Preview testing\n")
	defer t.Log("Preview testing finished\n")

	automata := InitialStateAutomaton1D()
	timeLaps := logic.NewTimeLaps(&automata)
	timeLaps.Next()

	got := (*timeLaps.Previous()).GetBoard().([]bool)
	expected := []bool{true, false, true, false, false}
	AutomatonCorrectnessChecker(t, got, expected)
}
