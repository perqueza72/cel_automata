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

func initCell2DOn(row, col int) ICell {
	return initCell2D(row, col, true)
}

func initPattern2D(positions [][]int) IPattern {

	cells := make([]*ICell, len(positions))
	for i, p := range positions {
		cell := initCell2DOn(p[0], p[1])
		cells[i] = &cell
	}

	pattern := logic.NewPattern(cells)

	return pattern
}

func initAutomata2D(pattern IPattern, board [][]bool) IAutomataCellular {
	cells := make([][]ICell, len(board))

	for i := 0; i < len(cells); i++ {

		row := make([]ICell, len(board))
		for j := range row {
			cell := initCell2D(i, j, board[i][j])
			row[j] = cell
		}
		cells[i] = row
	}

	return automatas.NewAutomata2D(cells, &pattern, 0)
}

func InitialState() (IAutomataCellular, [][]bool) {
	positions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	board := [][]bool{
		{false, true, false, true, false},
		{true, false, true, false, true},
		{false, true, false, true, false},
		{true, false, true, false, true},
		{false, true, false, true, false}}
	pattern := initPattern2D(positions)
	automata := initAutomata2D(pattern, board)

	return automata, board
}

func PossibleStates() [][][]bool {
	states := make([][][]bool, 2)
	states[0] = [][]bool{
		{false, true, false, true, false},
		{true, false, true, false, true},
		{false, true, false, true, false},
		{true, false, true, false, true},
		{false, true, false, true, false}}

	states[1] = [][]bool{
		{true, false, true, false, true},
		{false, true, false, true, false},
		{true, false, true, false, true},
		{false, true, false, true, false},
		{true, false, true, false, true},
	}

	return states
}

func TestAutomaton2D(t *testing.T) {

	automata, board := InitialState()
	states := PossibleStates()

	t.Logf("Initial board: %v\n", automata.GetBoard().([][]bool))
	Automaton2DCorrectnessChecker(t, automata.GetBoard().([][]bool), board)

	new_automata := automata.Transition()

	got := (*new_automata).GetBoard().([][]bool)
	expected := states[1]

	Automaton2DCorrectnessChecker(t, got, expected)
}

func Automaton2DCorrectnessChecker(t *testing.T, got [][]bool, expected [][]bool) bool {
	work := true

	for i := range got {
		for j := range got[i] {

			if got[i][j] != expected[i][j] {
				t.Errorf("In index %v got %v, whanted %v\n", i, got[i][j], expected[i][j])
				t.Fail()
				work = false
			}
		}
	}

	return work
}

func TestAutomaton2DTimeLapsNext(t *testing.T) {
	automata, _ := InitialState()
	states := PossibleStates()

	timeLaps := logic.NewTimeLaps(&automata)
	got := (*timeLaps.Next()).GetBoard().([][]bool)
	expected := states[1]

	Automaton2DCorrectnessChecker(t, got, expected)

	got = (*timeLaps.Next()).GetBoard().([][]bool)
	expected = states[0]
	Automaton2DCorrectnessChecker(t, got, expected)

	got = (*timeLaps.Next()).GetBoard().([][]bool)
	expected = states[1]
	Automaton2DCorrectnessChecker(t, got, expected)

}

func TestAutomaton2DTimeLapsGet(t *testing.T) {
	automata, _ := InitialState()
	timeLaps := logic.NewTimeLaps(&automata)
	states := PossibleStates()

	ITERATIONS := 100

	for i := 0; i < ITERATIONS; i++ {
		timeLaps.Next()
	}

	fault_cases := []int{}

	for i := 0; i <= ITERATIONS; i++ {
		got := (*timeLaps.Get(uint(i))).GetBoard().([][]bool)
		expected := states[i%2]
		if Automaton2DCorrectnessChecker(t, got, expected) {
			fault_cases = append(fault_cases, i)
		}
	}

	t.Logf("\n\nFault_cases: %v\n", fault_cases)
}

func TestAutomaton2DTimeLapsPreview(t *testing.T) {
	t.Log("Preview testing\n")
	defer t.Log("Preview testing finished\n")

	automata, _ := InitialState()
	states := PossibleStates()
	timeLaps := logic.NewTimeLaps(&automata)
	timeLaps.Next()

	got := (*timeLaps.Previous()).GetBoard().([][]bool)
	expected := states[0]
	Automaton2DCorrectnessChecker(t, got, expected)
}
