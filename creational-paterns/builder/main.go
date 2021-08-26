package builder

func Run() {
	kar98Builder := getBuilder("Kar98")
	bazookaBuilder := getBuilder("Bazooka")

	director := newGunDirector(kar98Builder)
	kar98Gun := director.buildGun()
	kar98Gun.Shot()

	director.setBuilder(bazookaBuilder)
	bazookaGun := director.buildGun()
	bazookaGun.Shot()
}
