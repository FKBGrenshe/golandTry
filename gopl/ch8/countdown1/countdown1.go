package countdown1

import (
	"fmt"
	"time"
)

func Countdown() {
	fmt.Println("commencing countdown")
	tick := time.Tick(1 * time.Second)
	for countdown := 3; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	Launch()
}

func Launch() {
	fmt.Println("Lisft off!")
}
