package abstract_factory

func Run()  {
	vnFactory := getWeaponFactory("VN")
	gunVN := vnFactory.makeGun()
	gunVN.shot()
	bomVN := vnFactory.makeBom()
	bomVN.explode()

	cnFactory := getWeaponFactory("CN")
	gunCN := cnFactory.makeGun()
	gunCN.shot()
	bomCN := cnFactory.makeBom()
	bomCN.explode()
}