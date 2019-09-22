package main

import (
	"fmt"
	"time"
)

// Retorna um channel somente leitura
func falar(pessoa string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d", pessoa, i)
		}

		fmt.Println("falar morreu")
	}()

	return c
}

// Multiplexar: junta dois channels e retorna um
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			fmt.Println("Travando ...")
			select {
			// o select é bloqueante, fica aguardando mesmo que as goroutines tinham morrido
			case s := <-entrada1:
				c <- s
			case s := <-entrada2:
				c <- s
			}
			fmt.Println("Rodando...")
		}
	}()

	return c
}

func main() {
	c := juntar(falar("João"), falar("Maria"))
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	time.Sleep(time.Second * 5)
}
