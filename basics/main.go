package main

import (
	"fmt"
	"time"
	"sync"
)
wg := sync.WaitGroup()
func main() {
	fmt.Println("Start")
	go process()
	time.Sleep(time.Millisecond * 10) // this is bad don't do it
	fmt.Println("done")
}

func process() {
	fmt.Println("processing")
}
