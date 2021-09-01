package decorator

type Accessory struct {
	Gun     IGun
	DameAdd int
}

func (a *Accessory) GetDame() int {
	if a.Gun != nil {
		return a.Gun.GetDame() + a.DameAdd
	}
	return 0
}

func (a *Accessory) getDameAfterAddItem() int {
	return a.Gun.GetDame() + a.DameAdd
}
