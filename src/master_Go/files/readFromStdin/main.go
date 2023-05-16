// Считывание пользовательского ввода из стандартного ввода Stdin
package main

import (
	"bufio"
	"fmt"
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
}
