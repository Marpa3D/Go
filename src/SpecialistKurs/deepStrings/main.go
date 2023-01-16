// Углубленное изучение типа string
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Пример строки"
	fmt.Println(str)
	// 1. Строка - это срез байт, со своими особенностями при обращении
	// к нижележащему массиву

	word := "Тестовая строка"
	fmt.Printf("Строка (string): %s, Длинна: %d\n", word, len(word))

	// Какие значения байт сейчас находятся в срезе word?
	fmt.Printf("Байты: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%x ", word[i]) // %c - это формат(флаг) представления символа
	}
	fmt.Println()

	// Каким образом получить доступ к отдельным символам?
	fmt.Printf("Символы: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c ", word[i]) // %x - это формат представления (глагол) 16-тиричного байта
	}
	fmt.Println()
	// 2. Строки в Go хранятся в виде набора UTF-8 символов.
	// Каждый символ может знаимать от 1 до 4-х байт!

	// 3. Руна (Rune) - это встроенный тип в Go, псевдоним (alias) типа int32,
	//  позволяет хранить ЕДИНЫЙ, НЕДЕЛИМЫЙ символ в UTF-8, вне зависимости, сколько байт он занимает!
	fmt.Printf("Руны: ")
	rSlice := []rune(word) // Преобразование среза байт в срез рун
	fmt.Printf("Len: %d Cap: %d\n", len(rSlice), cap(rSlice))
	for i := 0; i < len(rSlice); i++ {
		fmt.Printf("%c ", rSlice[i]) // %x - это формат представления (глагол) 16-тиричного байта
	}
	fmt.Println()

	// 4. Итерирование по строке с использованием рун
	for id, val := range word { // id - это индекс байта, с КОТОРОГО НАЧИНАЕТСЯ РУНА (val)!
		fmt.Printf("%c начинается с позиции: %d\n", val, id)
	}

	// 5. Создание строки из среза (слайса) байт
	mySliceByte := []byte{0x41, 0x42, 0x43, 0x44} // Можно испльзовать и десятичные значения, т.к. byte - это uint8
	mySliceString := string(mySliceByte)
	fmt.Printf("%s\n", mySliceString)

	// 6. Создание строки из слайса(среза) рун
	// руны как HEX
	myRuneHexSlice := []rune{0x41, 0x42, 0x43, 0x44}
	myHexSlice := string(myRuneHexSlice)
	fmt.Println(myHexSlice)

	// руны как литералы
	myRuneSlice := []rune{'h', 'e', 'l', 'l', 'o'}
	myStrSlice := string(myRuneSlice)
	fmt.Println(myStrSlice)
	fmt.Printf("%s это %T\n", myStrSlice, myStrSlice)

	// 7. Длинна и емкость строки
	// Длинна len() - это колличество байт в срезе!!!
	fmt.Println("Длина слова Марк:", len("Марк"), "bytes") // Длинна 8 байт

	// Длинна runeCounter - это колличество рун (символов)
	fmt.Println("Длина runes в слове Марк:", utf8.RuneCountInString("Марк"), "runes")

	// Вычисление емкости строки бессмысленно, т.к. строка(string) - это базовый тип в Go!

}
