package adapter

import "fmt"

type AdapterScopeAKM struct {
	AKMGun *AKM
}

func (a AdapterScopeAKM) ZoomWithScopeX8() {
	fmt.Println("Add scope X4 before X2")
	a.AKMGun.ZoomWithScopeX2()
	fmt.Println("Convert scope X8 from X2 done!")
}
