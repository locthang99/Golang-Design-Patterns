package proxy

import "fmt"

type GunSystem struct {
	Guns          []IShotHandle
	IndexShotGun  int
	MaxShot       int
	GunSystemName string
}

func (gs *GunSystem) HandleShotRequest(boss string) {
	fmt.Println("Gun System: ", gs.GunSystemName, " accept request from boss: ", boss)
	if gs.MaxShot > 0 {
		idx := gs.GetIndexShotGun()
		gs.Guns[idx].HandleShotRequest(gs.GunSystemName)
		gs.MaxShot--
	} else {
		fmt.Println("Can't shot because max shot = 0")
	}

}

func (gs *GunSystem) GetIndexShotGun() int {
	if gs.IndexShotGun >= len(gs.Guns)-1 {
		gs.IndexShotGun = 0
		return 0
	} else {
		gs.IndexShotGun++
		return gs.IndexShotGun
	}
}
