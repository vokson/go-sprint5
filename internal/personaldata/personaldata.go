package personaldata

import "fmt"

const personalPrintFormat = `Имя: %s
Вес: %.2f
Рост: %.2f
`

type Personal struct {
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	fmt.Printf(personalPrintFormat, p.Name, p.Weight, p.Height)
}
