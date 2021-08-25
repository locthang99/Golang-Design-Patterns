package abstract_factory

import "fmt"

type Gun struct {
	Name string
	Power int
}

func (g Gun) shot()  {
	fmt.Println(g.Name," Shot with dame: ",g.Power)
}