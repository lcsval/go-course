package main

import (
	"fmt"
	"time"

	"github.com/cod3rcursos/html"
)

func oMaisrapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	// estrutura de controle específica para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos"
		// default:
		// 	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisrapido(
		"https://www.google.com.br",
		"https://www.amazon.com.br",
		"https://www.youtube.com.br",
	)
	fmt.Println(campeao)
}
