package proxy

import "fmt"

type Gun struct {
	Name string
	Dame int
}

func (g *Gun) HandleShotRequest(boss string) {
	fmt.Println("Gun: ", g.Name, " accept request from boss: ", boss)
	fmt.Println("Shot with dame: ", g.Dame)
}
