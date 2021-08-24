package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}
type ObjectA struct {

}
var instanceA *ObjectA

func (o ObjectA) getAInstance() *ObjectA{
	if instanceA == nil{
		lock.Lock()
		defer lock.Unlock()
		if instanceA == nil{
			fmt.Println("Create instance A")
			instanceA = &ObjectA{}
		} else {
			fmt.Println("Instance A already created")
		}
	} else {
		fmt.Println("Instance A already created")
	}
	return instanceA
}