package main

import (
	"github.com/urjitbhatia/cozgo"
)

func cozDemo() {
	cozgo.Begin("mainstart")
	defer cozgo.End("mainstart")

	for i := 0; i < (100000 * 100000); i++ {
		if i%200 == 0 {
			cozgo.NamedProgress("mod200")
		} else if i%100 == 0 {
			cozgo.NamedProgress("mod100")
		}
	}
}

func main() {
	cozDemo()
}
