package decorator

import "fmt"

type Gun struct {
	Name string
	Dame int
}

func (g *Gun) GetDame() int {
	return g.Dame
}

func (g Gun) Shot() {
	fmt.Printf("Gun: %s shot with dame %d\n", g.Name, g.Dame)
}
