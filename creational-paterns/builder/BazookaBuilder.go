package builder

type BazookaBuilder struct {
	Name   string
	Scope  string
	Barrel string
	Bullet int
}

func newBazookaBuilder() *BazookaBuilder {
	return &BazookaBuilder{}
}

func (b *BazookaBuilder) setName() {
	b.Name = "Bazooka"
}

func (b *BazookaBuilder) setScope() {
	b.Scope = "X2"
}

func (b *BazookaBuilder) setBarrel() {
	b.Barrel = "Non"
}

func (b *BazookaBuilder) setBullet() {
	b.Bullet = 1
}

func (b *BazookaBuilder) getGun() Gun {
	return Gun{
		Name:   b.Name,
		Scope:  b.Scope,
		Barrel: b.Barrel,
		Bullet: b.Bullet,
	}
}
