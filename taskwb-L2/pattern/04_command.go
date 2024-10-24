package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*

	Применимость:
		1. 	Когда нужно параметризовать объекты выполняемым действием.
		2.	Когда нужно ставить операции в очередь, выполнять их по расписанию или передавать по сети.
		3.	Когда нужна операция отмены.

	Плюсы и минусы:
		+ 	Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
		+	Позволяет реализовать простую отмену и повтор операций.
		+	Позволяет реализовать отложенный запуск операций.
		+	Позволяет собирать сложные команды из простых.
		+	Реализует принцип открытости/закрытости.
		-	Усложняет код программы из-за введения множества дополнительных классов.

*/

func main() {
	game := Game{}
	svc := SaveCommand{
		Game:    &game,
		Name:    "Thrall",
		Species: "orc",
		Age:     56,
	}
	createButton := CreateButton{svc}
}

type Command interface {
	execute()
}

type SaveCommand struct {
	Game *Game
	// params:
	Name    string
	Species string
	Age     int
}

func (sc SaveCommand) execute() {
	sc.Game.CreateCharacter(sc.Name, sc.Species, sc.Age)
}

type CreateButton struct {
	Command
}

type Game struct {
	//
}

func (g *Game) CreateCharacter(name, species string, age int) {
	// logic
}
