package test

import (
	automatas "automata_automatas"
	logic "automata_logic"
	. "own_interfaces"
	"testing"
)

func initCell2D(row, col int, state bool) ICell {

	position := logic.NewPosition2D(row, col)
	cell := logic.NewCell(state, *position)
	return cell
}

func initPattern2D() IPattern {

	a := initCell2D(-1, 0, true)
	b := initCell2D(1, 0, true)

	cells := make([]*ICell, 2)
	cells[0] = &a
	cells[1] = &b
	pattern := logic.NewPattern(cells)

	return pattern
}

func initAutomata2D() IAutomataCellular {
	states := [5]bool{true, false, true, false, false}

	cells := make([][]ICell, 5)

	for i := 0; i < len(states); i++ {
		row := make([]ICell, 5)
		for j := range row {
			cell := initCell2D(i, j, states[j])
			row[j] = cell
		}
		cells[i] = row
	}

	pattern := initPattern2D()
	return automatas.NewAutomata2D(cells, &pattern, 0)
}

func TestAutomaton2D(t *testing.T) {

	automata := initAutomata2D()
	t.Log(automata.GetBoard())
	new_automata := automata.Transition()

	got := (*new_automata).GetBoard().([][]bool)
	expected := [][]bool{
		{false, true, false, false, false},
		{false, true, false, false, false},
		{false, true, false, false, false},
		{false, true, false, false, false},
		{false, true, false, false, false}}

	Automaton2DCorrectnessChecker(t, got, expected)
}

func Automaton2DCorrectnessChecker(t *testing.T, got [][]bool, expected [][]bool) {
	for i := range got {
		for j := range got[i] {

			if got[i][j] != expected[i][j] {
				t.Errorf("In index %v got %v, whanted %v\n", i, got[i][j], expected[i][j])
				t.Fail()
			}
		}
	}
}

func TestAutomaton2DTimeLapsNext(t *testing.T) {
	automata := initAutomata()

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

func TestAutomaton2DTimeLapsGet(t *testing.T) {
	automata := initAutomata()

	timeLaps := logic.NewTimeLaps(&automata)

	for i := 0; i < 100; i++ {
		timeLaps.Next()
		t.Log(len(timeLaps.Automatas))
	}

	got := (*timeLaps.Get(0)).GetBoard().([]bool)
	expected := []bool{true, false, true, false, false}
	t.Log("In id 0\n")
	AutomatonCorrectnessChecker(t, got, expected)

	got = (*timeLaps.Get(1)).GetBoard().([]bool)
	expected = []bool{false, true, false, false, false}
	t.Log("In id 1\n")
	AutomatonCorrectnessChecker(t, got, expected)

	got = (*timeLaps.Get(2)).GetBoard().([]bool)
	expected = []bool{true, false, false, false, false}
	t.Log("In id 2\n")
	AutomatonCorrectnessChecker(t, got, expected)

	got = (*timeLaps.Get(3)).GetBoard().([]bool)
	expected = []bool{false, false, false, false, false}
	t.Log("In id 3\n")
	AutomatonCorrectnessChecker(t, got, expected)

	got = (*timeLaps.Get(70)).GetBoard().([]bool)
	expected = []bool{false, false, false, false, false}
	t.Log("In id 70\n")
	AutomatonCorrectnessChecker(t, got, expected)
}

func TestAutomaton2DTimeLapsPreview(t *testing.T) {
	t.Log("Preview testing\n")
	defer t.Log("Preview testing finished\n")

	automata := initAutomata()
	timeLaps := logic.NewTimeLaps(&automata)
	timeLaps.Next()

	got := (*timeLaps.Previous()).GetBoard().([]bool)
	expected := []bool{true, false, true, false, false}
	AutomatonCorrectnessChecker(t, got, expected)
}
