// Практика по типу string
package main

import "fmt"

func main() {
	// Создание строки через краткую нотацию
	str := "Hello! Это строка на кирилице!" // И string не изменяема
	fmt.Println(str)
	fmt.Printf("Тип данных: %T\n", str)

	// Строка - это срез байт ("фрагмент, отрезок"), а тип byte - this uint8, поэтому
	bs := []byte(str)
	fmt.Println(bs)                    // выводит коды символов из UTF - 8
	fmt.Printf("Тип данных: %T\n", bs) // type []uint8
	bs = append([]byte(str[:2]), 'O', 'O', 'U')
	fmt.Printf("Добавлен символ: %q\n  для строки %s", bs, str)
	// Итерация по строке
	for i := 0; i < len(str); i++ {
		fmt.Printf("%#U\n ", str[i])
	}

	for i, val := range str {
		fmt.Printf("Индекс позиции %d, в HEX: %#x, И символ: %q\n ", i, val, val)
	}
}
