package facade

import "fmt"

type Bullet struct {
	Name  string
	Dame  int
	Total int
}

func (b *Bullet) Fly() {
	if b.Total > 0 {
		b.Total--
		fmt.Printf("Bullet %s fly with dame: %d\n", b.Name, b.Dame)
	} else {
		fmt.Println("Total bullet = 0 so can't fly")
	}
}
