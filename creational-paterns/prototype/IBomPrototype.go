package prototype

type IBomPrototype interface {
	explode() int
	clone() IBomPrototype
}
