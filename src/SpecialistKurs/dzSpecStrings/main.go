//На вход программе подается строка.
//Если из первого и последнего символов этой строки можно собрать слово "да" ("Да", "дА", "ДА", "да")
//- то программа выводит "СОГЛАСЕН" и "НЕ СОГЛАСЕН" в противном случае.
package main

import (
	"fmt"
)

func main() {
	str := "ДА, проверкА"
	result := []rune(str)
	//resRune := utf8.RuneCountInString(str)
	resItog := string(result[0]) + string(result[len(result)-1])
	if resItog == "да" || resItog == "ДА" || resItog == "дА" || resItog == "Да" {
		fmt.Println("СОГЛАСЕН")
	} else {
		fmt.Println("НЕ СОГЛАСЕН")
	}
}
