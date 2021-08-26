package builder

type Kar98Builder struct {
	Name   string
	Scope  string
	Barrel string
	Bullet int
}

func newKar98Builder() *Kar98Builder {
	return &Kar98Builder{}
}

func (k *Kar98Builder) setName() {
	k.Name = "Kar98"
}

func (k *Kar98Builder) setScope() {
	k.Scope = "X8"
}

func (k *Kar98Builder) setBarrel() {
	k.Barrel = "Silence"
}

func (k *Kar98Builder) setBullet() {
	k.Bullet = 5
}

func (k *Kar98Builder) getGun() Gun {
	return Gun{
		Name:   k.Name,
		Scope:  k.Scope,
		Barrel: k.Barrel,
		Bullet: k.Bullet,
	}
}
