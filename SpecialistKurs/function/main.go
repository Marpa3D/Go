package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])
	res1 := n1 + n2
	res2 := n1 * n2

	fmt.Println("Сумма аргументов:", res1, "Произведение аргументов:", res2)
}
