// программа проверяет, сдан ли экзамен
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Введите значение баллов:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	if grade > 60 {
		fmt.Println("Вы сдали экзамен!")
	} else {
		fmt.Println("Вам нужно еще подготовиться!..")
	}

}
