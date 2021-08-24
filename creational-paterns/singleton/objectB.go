package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type ObjectB struct {

}
var instanceB *ObjectB

func (o ObjectB) getBInstance()  *ObjectB{
	if instanceB == nil{
		once.Do(
			func() {
				fmt.Println("Create instance B")
				instanceB = &ObjectB{}
			},
		)
	} else {
		fmt.Println("Instance B already created")
	}
	return instanceB
}

func (o ObjectB) doSomethingB()  {
	fmt.Println("Object B do something")
}