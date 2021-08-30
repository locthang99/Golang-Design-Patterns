package bridge

import "fmt"

type ScopeX8 struct {
}

func (s *ScopeX8) ZoomTargetOnScope() {
	fmt.Println("Scope X8 get target")
}
