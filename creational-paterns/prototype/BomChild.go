package prototype

import "fmt"

type BomChild struct {
	Weight int
}

func (b *BomChild) explode() int {
	fmt.Printf("Bom child explode with weigth: %d\n", b.Weight)
	w := b.Weight
	b.Weight = 0
	return w
}

func (b *BomChild) clone() IBomPrototype {
	return &BomChild{Weight: b.Weight}
}
