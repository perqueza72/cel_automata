package models

type Cell struct {
	state bool
}

func (cell *Cell) changeState() {
	cell.state = !cell.state
}

func (cell *Cell) Transition() {
	cell.changeState()
}
