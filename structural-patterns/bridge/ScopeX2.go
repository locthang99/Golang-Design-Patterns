package bridge

import "fmt"

type ScopeX2 struct {
}

func (s *ScopeX2) ZoomTargetOnScope() {
	fmt.Println("Scope X2 get target")
}
