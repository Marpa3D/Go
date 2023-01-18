// Считывание и вывод строк
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var name string
	input := bufio.NewScanner(os.Stdin)
	input.Scan() // Комманда захвата потока ввода (Stdin)
	// и сохранение в буфер (до символа окончания строки \n)
	if input.Scan() {
		name = input.Text() // Команда возвращения элементов в буфере (отдаст строку)
	}
	fmt.Println(name)

	fmt.Println("Цикл for loop стартует:")
	for {
		if input.Scan() {
			result := input.Text()
			if result == "" {
				break
			}
			fmt.Println("Input string is:", result)
		}
	}

	// Преобразование строкового литерала к чему-нибудь числовому
	numStr := "44"
	res, err := strconv.Atoi(numStr) // Atoi - Anysing to int (именно int)!!!
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d is type %T\n", res, res)

	// Если нужно строку преобразовать в конкретный тип int (например, int64)
	numInt64, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d %T\n", numInt64, numInt64)
	// для чисел с плавающей запятой
	numFloat32, _ := strconv.ParseFloat(numStr, 32) // возвращает 64 -е число, т.к.
	//  32 ГАРАНТИРОВАНО влезает в 64-е число!
	fmt.Printf("%0.2f %T\n", float32(numFloat32), float32(numFloat32))

	numFloat64, _ := strconv.ParseFloat(numStr, 32)
	fmt.Printf("%0.4f %T\n", numFloat64, numFloat64)
}
