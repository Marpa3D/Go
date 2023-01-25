// Курс Дерека Банаса. Закрепляем основы
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var pl = fmt.Println

func main() {
	pl("Как Ваше имя?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Привет,", name)
	} else {
		log.Fatal(err)
	}
	pl("====================================")

	// Некоторые явные преобразования типов
	v := "880000000000"
	res, err := strconv.Atoi(v) // Строка в целочисленный тип
	pl(res, err, reflect.TypeOf(res))

	v1 := 880000000000
	res1 := strconv.Itoa(v1) // из целого числа в строку
	pl(res1, reflect.TypeOf(res1))
	pl("=====================================")

	v3 := "3.14"
	if vFloat, err := strconv.ParseFloat(v3, 64); err == nil {
		pl(vFloat)
	}
	s := fmt.Sprintf("%0.2f", 3.14)
	pl(s)
	pl("=====================================")

	// Функции для работы со строками
	str := "Это пример строки!"
	replacer := strings.NewReplacer("Это", "Классный") // Замена "Это" на "Классный"
	resultReplace := replacer.Replace(str)
	pl(resultReplace)

	pl("Length:", len(str)) // колличество байт в строке, а не символов
	strRune := []rune(str)
	fmt.Printf("Длина (колличесево) символов(рун):%c\t %d\n", strRune, len(strRune)) // кол-во символов
	pl("Строка содержит это:", strings.Contains(str, "пример"))                      // проверяет, содериж ли это слово строка?

	// Узнать индекс символа в строке
	pl("Индекс буквы \"е\" в строке \"Это пример строки\"", strings.Index(str, "о"))

	// Замена символа и поиск его повсюду
	pl("Замена символа и поиск повсюду:", strings.Replace(resultReplace, "с", "0", -1))

	// Удаление пробелов, escape-последовательностей
	str1 := "\nНесколько-слов-для-теста\n"
	pl("Обрезаем лишнее:", strings.TrimSpace(str1))
	pl("Split \"-\": ", strings.Split(str1, "-"))
	pl("Lower register:", strings.ToLower(str1))
	pl("Upper register:", strings.ToUpper(str1))
	pl("Prefix:", strings.HasPrefix("суперкот", "супер"))    // проверка, есть ли такой префикс?
	println("Suffix:", strings.HasSuffix("суперкот", "кот")) // а суффикс? true || false
	pl("=====================================")
}
