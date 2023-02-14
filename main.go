package main

import (
	"fmt"
	"time"
)

var ch = make(chan map[string]struct{}) // blocks main

func main() {
	fmt.Println("main start...")

	go func() { // separate routine to wait 3 seconds and populate ch
		fmt.Println("start timer...")
		time.Sleep(3 * time.Second)
		ch <- map[string]struct{}{"": {}}
	}()

	<-ch // blocking but unblocks when receives
	fmt.Println("main end...")
}
