package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*

	Применимость:
		1. Необходимость изоляции клиента от сложной подсистемы, добавлением простого интерфейса для взаимодействия
		2. Разделение на слои сложной подсистемы

	Плюсы и минусы:
		+	Изоляция клиента от реализации
		-	Есть риск, что фасад может стать божественным объектом, за счет размытой ответственности.
			И тогда получиться, что вместо простого интерфейса будет громоздкий
			с множеством функций разного по семантике назначения.
			Например: будет содержать в себе функции для работы с аудио и видео, хотя логичнее разделить это на два фасада

	Применимость:
		В библиотеке zap есть логгер sugar, который скрывает от пользователя детали, за счет использования интерфейсов
*/

type Facade struct {
	srv *server
}

func NewFacade(srv *server) *Facade {
	return &Facade{
		srv: srv,
	}
}

func (f *Facade) CreateAccountFacade(login, passwd string) {
	a := NewAccount(login, passwd)
	a.CheckAccountConstraints()
	f.srv.AddAccount(&a)
}

type account struct {
	//
}

func NewAccount(login, passwd string) account {}
func (a *account) CheckAccountConstraints()   {}

type server struct {
	//
}

func (s *server) AddAccount(a *account) {}
