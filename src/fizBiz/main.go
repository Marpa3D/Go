// FizBuz, замер скорости
package main

import (
	"fmt"
	"runtime"
	"time"
)

func fb(n int) {
	fiz := n%3 == 0
	buz := n%5 == 0

	switch {
	case fiz && !buz:
		fmt.Println("Fiz")
	case !fiz && buz:
		fmt.Println("Buz")
	default:
		fmt.Println("Other (", n, ")")
	}
}

func main() {
	startTime := time.Now()

	for i := 0; i < 50000; i++ {
		fb(i)
	}

	fmt.Println(runtime.NumCPU(), runtime.Version())
	elapced := time.Since(startTime)
	fmt.Printf("50 000 итераций за %.4f секунд", float64(elapced.Nanoseconds())/100_000_000_000.0)
}
