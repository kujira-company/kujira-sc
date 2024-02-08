package bug

import (
	"log"
)

func Run(n, i int) {
	if n == 0 { // Noncompliant
		doSomething()
	} else {
		doSomething()
	}

	switch i { // Noncompliant
	case 1:
		doSomething()
	case 2:
		doSomething()
	case 3:
		doSomething()
	default:
		doSomething()
	}
}

func doSomething() {
	log.Println("do something")
}
