package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, x := range dataset {
		err := dp.Parse(x)
		if err != nil {
			log.Println(err)
			continue
		}

		s, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Print(s)
	}
}
