package own_interfaces

type IAutomataCellular interface {
	Transition() error
	GetCell(position IPosition) ICell
}
