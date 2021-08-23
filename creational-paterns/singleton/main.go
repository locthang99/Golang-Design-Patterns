package creational_pattern_singleton

func Run()  {
	var (
		firstObjA ObjectA
		secondObjA ObjectA
		thirdObjA ObjectA
	)
	firstObjA.getAInstance()
	secondObjA.getAInstance()
	thirdObjA.getAInstance()

	var (
		firstObjB ObjectB
		secondObjB ObjectB
		thirdObjB ObjectB
	)
	firstObjB.getBInstance()
	secondObjB.getBInstance()
	objB := thirdObjB.getBInstance()

	objB.doSomethingB()

}