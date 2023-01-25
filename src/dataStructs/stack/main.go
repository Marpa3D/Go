// Структуры данных. Стек
package main

import (
	"dataStructs/stack/stackFuncs"
	"fmt"
)

func main() {
	var st stackFuncs.Stack
	st.Push(10)
	fmt.Printf("%+v\t adress: %p\t Длинна в байтах: %d\n", st, &st, len(st.items))

	st.Push(20)
	st.Push(30)
	st.Push(40)

	fmt.Printf("%+v\t adress: %p\t Длинна в байтах: %d\n", st, &st, len(st.items))

	st.Pop()
	fmt.Printf("%+v\t adress: %p\t Длинна в байтах: %d\n", st, &st, len(st.items))

	st.Top()
	fmt.Println(st.Top())

	st.Print() // Выводит все, что есть в стеке
}
