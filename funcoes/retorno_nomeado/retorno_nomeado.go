package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return
}

func main() {
	fmt.Println(trocar(10, 20))
}
