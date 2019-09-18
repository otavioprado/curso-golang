package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { // não tem while
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	// laço infinito
	for {
		fmt.Printf("Para sempre...")
		time.Sleep(time.Second)
	}
}
