package prototype

import "fmt"

func Run() {
	bom1 := &BomChild{Weight: 1}
	bom2 := &BomChild{Weight: 20}
	mediumBom := &BomParent{Children: []IBomPrototype{bom1, bom2}}
	bigBom := &BomParent{Children: []IBomPrototype{mediumBom.clone(), mediumBom.clone()}}

	fmt.Println("Big bom explode:")
	bigBom.explode()
	fmt.Println("Medium bom explode:")
	mediumBom.explode()
	fmt.Println("Medium bom explode again:")
	mediumBom.explode()
}
