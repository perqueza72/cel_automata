package own_interfaces

type ICell interface {
	Transition() error
	GetPosition() IPosition
	GetState() bool
}

type IPosition interface {
	Append(position IPosition) IPosition
	GetPosition() interface{}
}

type IPattern interface {
	Check(automata IAutomataCellular, cell ICell) bool
}
