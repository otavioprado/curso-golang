package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("E agora?")
	ch <- 4 // Vai travar nessa linha até alguém consumir do channel
	fmt.Println("Executou!")
	ch <- 5
	ch <- 6
	ch <- 7
	ch <- 8
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second * 10)
	// fmt.Println(<-ch)
}
