package bridge

import "fmt"

type Gun struct {
	Scope IScope
	Name  string
	Dame  int
}

func (g *Gun) ZoomTarget() {
	fmt.Println(g.Name, " starting Zoom:")
	g.Scope.ZoomTargetOnScope()
}

func (g *Gun) SetScope(scope IScope) {
	g.Scope = scope
}
