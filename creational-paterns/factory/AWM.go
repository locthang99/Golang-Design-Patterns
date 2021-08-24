package factory

import "fmt"

type AWM struct {
	Gun
	Scope string
}

func (a AWM) shot()  {
	a.Gun.shot()
	fmt.Println("Scope: ",a.Scope)
}