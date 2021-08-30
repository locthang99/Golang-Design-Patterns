package bridge

func Run() {
	scopeX2 := &ScopeX2{}
	scopeX8 := &ScopeX8{}

	g1 := Gun{Dame: 100, Name: "AKM"}
	akm := AKM{Barrel: "Vip barrel"}
	akm.Gun = g1
	akm.SetScope(scopeX2)
	akm.ZoomTarget()
	akm.SetScope(scopeX8)
	akm.ZoomTarget()

	g2 := Gun{Dame: 100, Name: "AWM"}
	awm := AWM{}
	awm.Gun = g2
	awm.SetScope(scopeX8)
	awm.ZoomTarget()
}
