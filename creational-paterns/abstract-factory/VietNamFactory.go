package abstract_factory

type VietNamFactory struct{

}

func (v VietNamFactory) makeGun() IGun{
	return Gun{Name: "Gun VN",Power: 500}
}
func (v VietNamFactory) makeBom() IBom  {
	return Bom{Name: "B100 VN",Weight: 100}
}