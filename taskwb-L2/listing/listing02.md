Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1

В первом случае имеем именованный результат и defer обращается именно
к именованному результату и имеет возможность его изменить.

Во втором случае мы отдаём значение без привязки к имени переменной,
defer меняет х, но этот х остаётся в скоупе функции, а вызывающий получает
значение записанное в момент return.

defer'ы выполняются в обратном порядке их объявления (LIFO),
исполняются сразу после завершения функции, но до возвращения результатов вызывающему.
```
