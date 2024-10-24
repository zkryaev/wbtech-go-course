package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Применимость:
 1. Когда у вас есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния, причём типов состояний много, и их код часто меняется.
 2. Когда код класса содержит множество больших, похожих друг на друга, условных операторов, которые выбирают поведения в зависимости от текущих значений полей класса.
 3. Когда вы сознательно используете табличную машину состояний, построенную на условных операторах, но вынуждены мириться с дублированием кода для похожих состояний и переходов.

Плюсы и минусы:
  - Избавляет от множества больших условных операторов машины состояний.
  - Концентрирует в одном месте код, связанный с определённым состоянием.
  - Упрощает код контекста.
  - Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

type State interface {
	Lock(*Storage)
	Unlock(*Storage)
}

type LockState struct{}

func (s *LockState) Lock(storage *Storage) {
	fmt.Println("Storage is already locked.")
}

func (s *LockState) Unlock(storage *Storage) {
	fmt.Println("Unlocking storage.")
	storage.setState(&UnlockState{}) // Переход к состоянию разблокировки
}

type UnlockState struct{}

func (s *UnlockState) Lock(storage *Storage) {
	fmt.Println("Locking storage.")
	storage.setState(&LockState{}) // Переход к состоянию блокировки
}

func (s *UnlockState) Unlock(storage *Storage) {
	fmt.Println("Storage is already unlocked.")
}

type Storage struct {
	state State
}

func NewStorage() *Storage {
	// Изначально хранилище разблокировано
	return &Storage{state: &UnlockState{}}
}

func (s *Storage) setState(state State) {
	s.state = state
}

func (s *Storage) Lock() {
	s.state.Lock(s)
}

func (s *Storage) Unlock() {
	s.state.Unlock(s)
}

// Пример использования:
func main() {
	storage := NewStorage()

	storage.Lock()
	storage.Lock()
	storage.Unlock()
	storage.Unlock()
}
