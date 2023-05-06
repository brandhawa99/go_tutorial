package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
	}
}

func main() {
	fmt.Println("start")
	c := make(chan int)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}
	fmt.Println("middle")
	for {
		fmt.Println("Am here")
		c <- rand.Int()
		time.Sleep(time.Millisecond * 100)
	}
}
