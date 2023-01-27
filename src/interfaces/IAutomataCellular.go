package own_interfaces

type IAutomataCellular interface {
	Transition() *IAutomataCellular
	GetCell(position IPosition) (*ICell, bool)
	GetId() uint
	GetBoard() interface{}
	Copy() *IAutomataCellular
}

/***

This is pattern
1,2,1
1,x,1
1,1,1

Position, State
pattern = {((0,1),true)}



cell.position + pattron
*/
