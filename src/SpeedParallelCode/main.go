// Сравнение скорости выполнения при конкурентном
// и параллельном программировании
package main

import (
	"fmt"
	"runtime"
	"time"
)

// Две одиннаковые функции. Сравнение скорости выполнения при конкурентном
// и параллельном программировании
func main() {
	begin := time.Now()
	numCPU := runtime.NumCPU()
	fmt.Printf("Подключаем ядра процессора, шт.: %d\n", numCPU)
	runtime.GOMAXPROCS(4)
	go func() {
		for y := 0; y < 4; y++ {
			fmt.Println(y)
		}
	}()

	go func() {
		for y := 0; y < 4; y++ {
			fmt.Println(y)
		}
	}()
	elapsedTime := time.Since(begin)
	fmt.Println("Общее Время Выполнения:" + elapsedTime.String())
	time.Sleep(time.Second)
}
