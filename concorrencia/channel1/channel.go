package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // enviando dados para o canal (escrita)
	<-ch    // saindo do canal um determinado valor

	ch <- 2
	fmt.Println(<-ch)
}
