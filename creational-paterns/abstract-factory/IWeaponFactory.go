package abstract_factory

type IWeaponFactory interface {
	makeGun() IGun
	makeBom() IBom
}

func getWeaponFactory(country string) IWeaponFactory {
	if country=="VN"{
		return VietNamFactory{}
	} else {
		return ChinaFactory{}
	}
}