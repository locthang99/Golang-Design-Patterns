package facade

import "fmt"

type Scope struct {
	Name  string
	Ratio int
}

func (s *Scope) Zoom(ratio int) {
	s.Ratio = ratio
	fmt.Printf("Scope %s zoom object with ratio: %d\n", s.Name, s.Ratio)
}
