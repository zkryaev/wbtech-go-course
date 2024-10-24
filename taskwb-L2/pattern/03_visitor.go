package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*

	Применимость:
		- Когда вам нужно выполнить какую-то операцию над всеми элементами сложной структуры объектов, например, деревом.
		- Когда над объектами сложной структуры объектов надо выполнять некоторые не связанные между собой операции, но вы не хотите «засорять» классы такими операциями.
		- Когда новое поведение имеет смысл только для некоторых классов из существующей иерархии.

	Плюсы и минусы:
		+ Упрощает добавление операций, работающих со сложными структурами объектов.
		+ Объединяет родственные операции в одном классе.
		+ Посетитель может накапливать состояние при обходе структуры элементов.
		- Паттерн не оправдан, если иерархия элементов часто меняется.
		- Может привести к нарушению инкапсуляции элементов.
*/

type Visitor interface {
	MakeObj(Object)
	MakeAnotherObj(AnotherObject)
}

type Element interface {
	accept(v *Visitor)
}

type ConcreteVisitor struct {
}

func (cv *ConcreteVisitor) MakeObj(obj Object)               {}
func (cv *ConcreteVisitor) MakeAnotherObj(obj AnotherObject) {}

type Object struct {
	//
}

func (o *Object) Dance() {}

func (o *Object) accept(v Visitor) {
	v.MakeObj(*o)
}

type AnotherObject struct {
	//
}

func (ao *AnotherObject) StayCalm() {}

func (ao *AnotherObject) accept(v Visitor) {
	v.MakeAnotherObj(*ao)
}
