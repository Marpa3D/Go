// FizBuz, замер скорости
package main

import (
	"fmt"
	"sync"
	"time"
)

func fb(n int) {
	fiz := n%3 == 0
	buz := n%5 == 0

	switch {
	case fiz && !buz:
		fmt.Println("Fiz")
	case !fiz && buz:
		fmt.Println("FizBuz")
	default:
		fmt.Println("Other (", n, ")")
	}
}

func main() {
	startTime := time.Now()
	wg := sync.WaitGroup{}

	for i := 0; i < 50000; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			fb(i)
		}()
	}
	wg.Wait()
	elapced := time.Since(startTime)
	fmt.Printf("50 000 итераций за %.4f секунд\n", float64(elapced.Nanoseconds())/100_000_000_000.0)

	fmt.Printf("Done!")
}
