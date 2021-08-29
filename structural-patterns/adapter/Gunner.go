package adapter

import "fmt"

type Gunner struct {
	Name string
}

func (g *Gunner) ZoomWithGunVip(gun IGunVip) {
	fmt.Println("Gunner: ", g.Name)
	gun.ZoomWithScopeX8()
}
