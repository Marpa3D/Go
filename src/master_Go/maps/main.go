// Header map and cloning
package main

import "fmt"

func main() {
	pf := fmt.Printf
	// При объявлении карты Go создает указатель на Заголовок карты в памяти. Он ссылается на структуру
	// данны типа хэш-таблицы
	m1 := map[int]string{1: "Mark", 2: "Olivia"}
	pf("%v:%p\n", m1, &m1)

	// При копировании карты на другую карту, исходная(внутренняя) структура данных не меняется,
	// а просто ссылается. Обе карты имеют дин и тот же Заголовок карты
	m2 := m1
	m1[2] = "Maria"
	// у карт m1 и m2 - один Заголовок, ссылающийся на внутреннюю структуру данных
	pf("Начальная карта - %v: Конечная карта - %v\n", m1, m2) // Внутренняя структура данных та же!

}
