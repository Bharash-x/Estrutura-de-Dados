package main

import "fmt"

func main() {
	l := &Lista{}

	for _, v := range []int{10, 20, 50, 60, 80} {
		l.InserirFimSemTail(v)
	}
	fmt.Print("Inicial:                ")
	l.Imprimir()

	l.InserirInicio(5)
	fmt.Print("1) Inserir 5 no inicio: ")
	l.Imprimir()

	l.InserirFimSemTail(5)
	fmt.Print("2) Inserir 5 no fim:    ")
	l.Imprimir()

	v, ok := l.RemoverInicio()
	fmt.Printf("3) Removido head (%d):   ", v)
	_ = ok
	l.Imprimir()

	v, ok = l.RemoverFim()
	fmt.Printf("4) Removido tail (%d):   ", v)
	_ = ok
	l.Imprimir()
}
