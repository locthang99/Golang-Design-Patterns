package abstract_factory

import "fmt"

type Bom struct {
	Name string
	Weight int
}
func (b Bom) explode() {
	fmt.Println(b.Name," Explode with weight: ",b.Weight)
}