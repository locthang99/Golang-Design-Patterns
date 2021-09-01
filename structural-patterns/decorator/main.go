package decorator

import "fmt"

func Run() {
	gun := &Gun{Dame: 5, Name: "AKM"}
	gunScope := &Scope{}
	gunScope.DameAdd = 20
	gunScope.Gun = gun

	gunBarrelScope := &Barrel{Id: 1, Accessory: Accessory{DameAdd: 30}}
	gunBarrelScope.Gun = gunScope

	fmt.Println("Gun after wrap by scope and barrel with dame: ", gunBarrelScope.GetDame())
}
