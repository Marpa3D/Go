package main

import (
	"fmt"
)

func main() {
	var x1, b,c, d float64


	fmt.Scan(&x1, &b, &c, &d)

	if(x1 + 2 == c && b + 1 == d || x1 + 1 == c && b + 2 == d || x1 - 1 == c && b + 2 == d || x1 - 2 == c && b + 1 == d || x1 - 2 == c && b - 1 == d || x1 - 1 == c && b - 2 == d || x1 + 1 == c && b - 2 == d || x1 + 2 == c && b - 1 == d ) {

		fmt.Println("ДА")
	} else {
		fmt.Println("НЕТ")
	}
}
