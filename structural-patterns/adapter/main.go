package adapter

func Run() {
	gunner := Gunner{Name: "Gunner Vip Pro"}
	awm := AWM{}
	gunner.ZoomWithGunVip(awm)

	akm := AKM{}
	adapterScopeAKM := AdapterScopeAKM{AKMGun: &akm}
	gunner.ZoomWithGunVip(adapterScopeAKM)
}
