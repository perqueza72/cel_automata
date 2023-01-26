package own_interfaces

type ICell interface {
	SetState(state bool)
	GetPosition() *IPosition
	GetState() bool
}

type IPosition interface {
	Translate(position *IPosition) IPosition
	GetPosition() interface{}
}

type IPattern interface {
	Check(automata IAutomataCellular, cell *ICell) bool
}
