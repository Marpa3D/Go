// На вход программе подается строка.
// Если из первого и последнего символов этой строки можно собрать слово "да" ("Да", "дА", "ДА", "да")
// - то программа выводит "СОГЛАСЕН" и "НЕ СОГЛАСЕН" в противном случае.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите Ваше слово:")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	//str := "ДпроверкА"
	str := input.Text()
	result := []rune(str)
	resItog := string(result[0]) + string(result[len(result)-1]) // Конкатенация строк
	if resItog == "да" || resItog == "ДА" || resItog == "дА" || resItog == "Да" {
		fmt.Println("СОГЛАСЕН")
	} else {
		fmt.Println("НЕ СОГЛАСЕН")
	}
}
