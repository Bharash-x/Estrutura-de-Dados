package Atividade_01_N2

import "fmt"

type No struct {
	valor int
	ant   *No
	prox  *No
}

type Lista struct {
	head *No
	tail *No
}

func (l *Lista) Imprimir() {
	fmt.Print("[ ")
	for n := l.head; n != nil; n = n.prox {
		fmt.Print(n.valor, " ")
	}
	fmt.Println("]")
}

//Inserir no começo

func (l *Lista) InserirInicio(v int) {
	novo := &No{valor: v}

	if l.head == nil {
		l.head = novo
		l.tail = novo
		return
	}

	novo.prox = l.head
	l.head.ant = novo
	l.head = novo
}

//Inserir no final SEM usar o tail

func (l *Lista) InserirFimSemTail(v int) {
	novo := &No{valor: v}

	if l.head == nil {
		l.head = novo
		l.tail = novo
		return
	}

	atual := l.head
	for atual.prox != nil {
		atual = atual.prox
	}

	atual.prox = novo
	novo.ant = atual
	l.tail = novo
}

//Remover o head

func (l *Lista) RemoverInicio() (int, bool) {
	if l.head == nil {
		return 0, false // lista vazia
	}

	removido := l.head
	valor := removido.valor

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return valor, true
	}

	l.head = l.head.prox
	l.head.ant = nil
	removido.prox = nil

	return valor, true
}

//Remover o tail

func (l *Lista) RemoverFim() (int, bool) {
	if l.tail == nil {
		return 0, false
	}

	removido := l.tail
	valor := removido.valor

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return valor, true
	}

	l.tail = l.tail.ant
	l.tail.prox = nil
	removido.ant = nil

	return valor, true
}
