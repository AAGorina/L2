package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы,
а также реальные примеры использования данного примера на практике.
*/
// интерфейс обработки запросов
type Handler interface {
	HandOver(handler Handler) Handler
	Operate(info int)
}

// конкретные обработчики
// имеет доступ к преемнику
// если способны обработать запрос то обрабатывают
type ConcreteHandler_1 struct {
	next Handler
}

type ConcreteHandler_2 struct {
	next Handler
}

type ConcreteHandler_3 struct {
	next Handler
}

// реализация задания преемников
func (h *ConcreteHandler_1) HandOver(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *ConcreteHandler_2) HandOver(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *ConcreteHandler_3) HandOver(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *ConcreteHandler_1) Operate(info int) {
	if info >= 100 && info < 200 {
		fmt.Println("Concrete Handler 1 handelled operation")
		fmt.Println(info, " -> ", info)
	} else if h.next != nil {
		h.next.Operate(info)
	} else {
		fmt.Println("Handellers unable to handle operation.")
	}
}

func (h *ConcreteHandler_2) Operate(info int) {
	if info < 50 {
		fmt.Println("Concrete Handler 2 handelled operation")
		fmt.Println(info, " -> ", info*info*info)
	} else if h.next != nil {
		h.next.Operate(info)
	} else {
		fmt.Println("Handellers unable to handle operation.")
	}
}

func (h *ConcreteHandler_3) Operate(info int) {
	if info >= 50 && info < 100 {
		fmt.Println("Concrete Handler 3 handelled operation")
		fmt.Println(info, " -> ", info*info)
	} else if h.next != nil {
		h.next.Operate(info)
	} else {
		fmt.Println("Handellers unable to handle operation.")
	}
}

func main() {
	h1 := &ConcreteHandler_1{}
	h2 := &ConcreteHandler_2{}
	h3 := &ConcreteHandler_3{}

	h1.HandOver(h2).HandOver(h3)

	h1.Operate(15)
	h1.Operate(68)
	h1.Operate(157)
	h1.Operate(372)
}
