package factory

import "fmt"

func Run()  {
	akm,_ := GetGun("AKM")
	akm.shot()

	awm,_ := GetGun("AWM")
	awm.shot()

	_,err:=GetGun("ShotGun")
	if err != nil{
		fmt.Println(err)
	}
}