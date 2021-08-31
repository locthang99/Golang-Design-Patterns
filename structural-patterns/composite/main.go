package composite

import "fmt"

func Run() {
	bom1 := &BomChild{Weight: 1}
	bom2 := &BomChild{Weight: 2}
	mediumBom1 := &BomParent{Children: []IBom{bom1, bom2}}

	bom3 := &BomChild{Weight: 3}
	bom4 := &BomChild{Weight: 4}
	mediumBom2 := &BomParent{Children: []IBom{bom3, bom4}}

	bigBom := &BomParent{Children: []IBom{mediumBom1, mediumBom2}}

	fmt.Println("Big bom explode:")
	bigBom.explode()
}
