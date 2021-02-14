package shared

import "fmt"

type Bau struct {
	Asd string
	Isd int
}

func (b Bau) Miao() {
	fmt.Println("Miao", b.Asd)
}
