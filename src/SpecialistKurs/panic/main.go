// практика для panic
package main

import "fmt"

func highlow(high, low int) {
	if high < low {
		fmt.Println("Panic!!!")
		panic("in highlow()...")
	}
	fmt.Println("Deferred: highlow(", high, ",", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")

	highlow(high, low+1)
}

func main() {
	highlow(4, 0)
	fmt.Println("Программа завершена!")
}
