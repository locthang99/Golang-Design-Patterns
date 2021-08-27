package prototype

import "fmt"

type BomParent struct {
	Children []IBomPrototype
}

func (b *BomParent) explode() int {
	totalWeight := 0
	for _, child := range b.Children {
		totalWeight += child.explode()
	}
	fmt.Printf("Bom parent explode with weigth: %d\n", totalWeight)
	return totalWeight
}

func (b *BomParent) clone() IBomPrototype {
	var listChild []IBomPrototype
	for _, child := range b.Children {
		listChild = append(listChild, child.clone())
	}
	return &BomParent{Children: listChild}
}
