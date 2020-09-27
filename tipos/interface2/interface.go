package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferarri struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferarri) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferarri{"F40", false, 0}
	carro1.ligarTurbo()

	var carro2 esportivo = &ferarri{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
