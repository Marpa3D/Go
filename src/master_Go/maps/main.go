// Header map and cloning
package main

import "fmt"

func main() {
	pf := fmt.Printf
	// При объявлении карты Go создает указатель на Заголовок карты в памяти. Он ссылается на структуру
	// данны типа хэш-таблицы
	m1 := map[int]string{1: "Mark", 2: "Olivia"}
	pf("%v:%p\n", m1, &m1)
}
