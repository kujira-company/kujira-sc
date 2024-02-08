package bug

import "log"

func Run(condition bool) {
	if condition {
	} else if condition { // Noncompliant
		doSomething()
	}
}

func doSomething() {
	log.Println("do something")
}
