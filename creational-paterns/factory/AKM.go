package factory

import "fmt"

type AKM struct {
	Gun
	Barrel string
}

func (a AKM) shot()  {
	a.Gun.shot()
	fmt.Println("Barrel: ",a.Barrel)
}