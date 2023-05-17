// Считывание пользовательского ввода из стандартного ввода Stdin
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%T\n", scanner)

	// Сканируй, санер!)
	scanner.Scan()

	text := scanner.Text()
	bytes := scanner.Bytes()

	fmt.Println("Введеный текст:", text)
	fmt.Println("Прочитанные байты:", bytes)

	// Считывание, пока не введена команда "exit"
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Вы ввели текст:", text)
		if text == "exit" {
			fmt.Println("Выход из сканирования...")
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
