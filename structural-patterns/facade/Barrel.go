package facade

import "fmt"

type Barrel struct {
	Name   string
	Recoil int
}

func (b *Barrel) ReduceRecoil(ratioReduce int) {
	b.Recoil -= ratioReduce
	if ratioReduce > 0 {
		fmt.Println("Barrel reduce recoil: ", b.Recoil)
	} else {
		fmt.Println("Barrel increase recoil: ", b.Recoil)
	}
}
