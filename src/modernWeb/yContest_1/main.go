package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var str1, str2, str3 string
	fmt.Fscan(os.Stdin, &str1)
	fmt.Fscan(os.Stdin, &str2)
	fmt.Fscan(os.Stdin, &str3)

	s1, _ := strconv.Atoi(str1)
	s2, _ := strconv.Atoi(str2)
	s3, _ := strconv.Atoi(str3)
	if str1 == "раз" || str1 == "один" && str2 == "два" && str3 == "три" {
		fmt.Println("ОК")
		return
	}
	if s1 == 1 && s2 == 2 && s3 == 3 {
		fmt.Println("ОК")
	} else {
		fmt.Println("НЕ ПРАВИЛЬНО")
	}
}
