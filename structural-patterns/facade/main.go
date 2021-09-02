package facade

func Run() {
	gun := newGun(1, 2, 2, 4)
	gun.Zoom(5)
	gun.Shot()
	gun.Shot()
	gun.Shot()
	gun.addBullet(1)
	gun.Shot()
}
