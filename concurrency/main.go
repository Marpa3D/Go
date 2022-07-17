// Погружение в параллельное программирование многоядерных систем
package main

import (
	"fmt"
	"time"
)

func somesings(s string) {
	fmt.Println(s)
}

func main() {
	go somesings("First soprogram!")
	time.Sleep(8 * time.Second)
	somesings("Second soprogram!")
}
