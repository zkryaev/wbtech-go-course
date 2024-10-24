package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*

	Применимость:
		1. Когда программа должна обрабатывать разнообразные запросы несколькими способами, но заранее неизвестно,
		   какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
		2. Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
		3. Когда набор объектов, способных обработать запрос, должен задаваться динамически.

	Плюсы и минусы:
		+	Уменьшает зависимость между клиентом и обработчиками.
		+	Реализует принцип единственной обязанности.
		+	Реализует принцип открытости/закрытости.
		-	Запрос может остаться никем не обработанным.

*/

func main() {
	h1 := &ValidationHandler{}
	h2 := &SimpleHandler{}
	h1.SetNext(h2)
	h1.Handle(52)
}

type Handler interface {
	SetNext(h Handler)
	Handle(req Request)
}

type Request int

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(nextHandler Handler) {
	h.next = nextHandler
}

func (h BaseHandler) Handle(req Request) {
	fmt.Println("Hello #1")
	if h.next != nil {
		h.next.Handle(req)
	}
}

type ValidationHandler struct {
	BaseHandler
}

func (h *ValidationHandler) Handle(req Request) {
	fmt.Println("Hello #2")
	if h.next != nil {
		h.next.Handle(req)
	}
}

type SimpleHandler struct {
	BaseHandler
}

func (h *SimpleHandler) Handle(req Request) {
	fmt.Println("Hello #3")
	if h.next != nil {
		h.next.Handle(req)
	}
}
