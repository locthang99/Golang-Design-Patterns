package proxy

import "fmt"

func Run() {
	BossName := "Biladen vip pro"

	g1 := &Gun{Name: "Gun1", Dame: 10}
	g2 := &Gun{Name: "Gun2", Dame: 20}
	gunSys := GunSystem{GunSystemName: "GS500", IndexShotGun: -1, MaxShot: 3, Guns: []IShotHandle{g1, g2}}

	for i := 0; i < 5; i++ {
		fmt.Printf("------------- Request %d ---------------\n", i)
		gunSys.HandleShotRequest(BossName)
	}
}
