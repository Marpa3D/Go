// изучение bufio, log & os packages
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Введите Ваше имя: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Привет,%s", input)
}
