// Линейные структуры данных - список (List)
package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Список - это последовательность елементов. Имеет переменчивую длину.
	//Это базовый тип контейнера.
	var inList list.List
	inList.PushBack(8)
	inList.PushBack(22)
	inList.PushBack(33)

	for elem := inList.Front(); elem != nil; elem = elem.Next() {
		fmt.Println(elem.Value)
	}
}
