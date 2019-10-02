package example

import (
	cozgo "github.com/urjitbhatia/coz-go"
)

func sampledRun() {
	cozgo.Begin("mainstart")
	defer cozgo.End("mainstart")
	for i := 0; i < 1000; i++ {
		if i%100 == 0 {
			cozgo.NamedProgress("foo")
		}
		if i%200 == 0 {
			cozgo.NamedProgress("foo")
		}
	}
}
