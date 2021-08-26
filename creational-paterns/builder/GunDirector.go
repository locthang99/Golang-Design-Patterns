package builder

type GunDirector struct {
	builder IGunBuilder
}

func newGunDirector(b IGunBuilder) *GunDirector {
	return &GunDirector{
		builder: b,
	}
}

func (d *GunDirector) setBuilder(b IGunBuilder) {
	d.builder = b
}

func (d *GunDirector) buildGun() Gun {
	d.builder.setName()
	d.builder.setScope()
	d.builder.setBarrel()
	d.builder.setBullet()
	return d.builder.getGun()
}
