package abstract_factory

type ChinaFactory struct{

}

func (c ChinaFactory) makeGun() IGun{
	return Gun{Name: "Gun China",Power: 10}
}
func (c ChinaFactory) makeBom() IBom  {
	return Bom{Name: "B20 China",Weight: 20}
}
