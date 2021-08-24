package factory

import "fmt"

func GetGun(nameGun string) (IGun,error)  {
	if nameGun == "AKM"{
		return AKM{Gun:Gun{Name: "AKM",Power: 50},Barrel: "Silence"},nil
	}
	if nameGun == "AWM" {
		return AWM{Gun:Gun{Name: "AWM",Power: 100},Scope: "X8"},nil
	}
	return  nil,fmt.Errorf("Name Gun wrong")
}
