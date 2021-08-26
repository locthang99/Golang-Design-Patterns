package builder

import "fmt"

type Gun struct {
	Name   string
	Scope  string
	Barrel string
	Bullet int
}

func (g Gun) Shot() {
	fmt.Printf(" %s have: Scope: %s , Barrel: %s, Bullet: %d\n", g.Name, g.Scope, g.Barrel, g.Bullet)
}
