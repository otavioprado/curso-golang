package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

// Google I/O 2012 - Go Concunrrency Patterns
// <-chan - canal somente-leitura

// Generator Pattern:
func titulo(urls ...string) <-chan string {
	c := make(chan string, 1)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1] // Operação bloqueante
			fmt.Println("Só vai imprimir isso aqui uma vez pra cada channel!")
		}(url)
	}

	return c
}

func main() {
	t1 := titulo("https://google.com.br", "https://youtube.com.br") // cria um channel
	//t2 := titulo("https://amazon.com.br", "https://uol.com.br")     // cria outro channel
	fmt.Println(t1)

	time.Sleep(time.Second * 20)
	// fmt.Println("Primeiros:", <-t1, "|", <-t2)
	// fmt.Println("Segundos:", <-t1, "|", <-t2)
}

// Por default, o tamanho do buffer do channel (quando não informado) é zero.
// Quando o channel é unbuffered, toda operação de leitura / escrita é blocking, ou seja,
// só será possível postar no channel após alguém estar esperando por uma leitura.
