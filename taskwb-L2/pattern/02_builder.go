package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*

	Применимость:
		1. Избавиться от «телескопического конструктора» (ЯП где есть перегрузка методов)
		2. Когда ваш код должен создавать разные представления какого-то объекта. Например, деревянные и железобетонные дома.
		3. Когда вам нужно собирать сложные составные объекты, например, деревья (компоновщик)

	Плюсы и минусы:
		+ Позволяет создавать продукты пошагово
		+ Позволяет использовать один и тот же код для создания различных продуктов.
		+ Изолирует сложный код сборки продукта от его основной бизнес-логики.
		- Усложняет код программы из-за введения дополнительных классов.
		- Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.


*/

func main() {
	b := ConcreteBuilder{}
	d := Director{&b}
	d.Cook()
	_ = b.GetResult()
}

type Builder interface {
	Foundation()
	Walls()
	Doors()
	Windows()
}

type House struct {
}

type Director struct {
	b Builder
}

func (d *Director) Cook() {
	b.Foundation()
	b.Walls()
	b.Doors()
	b.Windows()
}

type ConcreteBuilder struct {
	h *House
}

func (cb *ConcreteBuilder) Foundation()      {}
func (cb *ConcreteBuilder) Walls()           {}
func (cb *ConcreteBuilder) Doors()           {}
func (cb *ConcreteBuilder) Windows()         {}
func (cb *ConcreteBuilder) GetResult() House { return *cb.h }
