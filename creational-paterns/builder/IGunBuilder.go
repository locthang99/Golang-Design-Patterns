package builder

type IGunBuilder interface {
	setName()
	setScope()
	setBarrel()
	setBullet()
	getGun() Gun
}

func getBuilder(builderName string) IGunBuilder {
	if builderName == "Kar98" {
		return newKar98Builder()
	}
	if builderName == "Bazooka" {
		return newBazookaBuilder()
	}
	return nil
}
