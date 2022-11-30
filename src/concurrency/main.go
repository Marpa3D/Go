// работа с типом группа горутин (Wait Group)
package main

import (
	"fmt"
	"sync"
)

func somesings(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {

	var wg sync.WaitGroup

	words := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
	}
	wg.Add(8)

	for i, x := range words {
		go somesings(fmt.Sprintf("%d: %s", i, x), &wg)
	}
	wg.Wait()
	wg.Add(1)
	//go somesings("First soprogram!")
	// time.Sleep(8 * time.Second) не используем такой пример в "лоб"
	somesings("Second soprogram!", &wg)
}
