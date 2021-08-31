package composite

import "fmt"

type BomParent struct {
	Children []IBom
}

func (b *BomParent) explode() int {
	totalWeight := 0
	for _, child := range b.Children {
		totalWeight += child.explode()
	}
	fmt.Printf("Bom parent explode with weigth: %d\n", totalWeight)
	return totalWeight
}
