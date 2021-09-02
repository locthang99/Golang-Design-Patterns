package facade

type Gun struct {
	scope  *Scope
	barrel *Barrel
	bullet *Bullet
}

func newGun(ratio int, recoil int, totalBullet int, dame int) *Gun {
	g := &Gun{}
	g.scope = &Scope{Ratio: ratio}
	g.barrel = &Barrel{Recoil: recoil}
	g.bullet = &Bullet{Total: totalBullet, Dame: dame}
	return g
}

func (g *Gun) Zoom(ratio int) {
	g.scope.Zoom(ratio)
}

func (g *Gun) AddBarrel(recoil int) {
	g.scope.Zoom(1 + recoil)
	g.barrel.ReduceRecoil(recoil)
}

func (g *Gun) Shot() {
	g.barrel.ReduceRecoil(-2 * g.barrel.Recoil)
	g.bullet.Fly()
}

func (g *Gun) addBullet(n int) {
	g.bullet.Total += n
}
