// Работа со временем. Пакет "time"
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pl = fmt.Println

func main() {
	now := time.Now()
	pl(now.Day(), now.Month(), now.Year())
	pl(now.Hour(), now.Minute(), now.Second())

	// Генерация случайного числа
	tNow := time.Now().Unix()
	rand.Seed(tNow)
	randNum := rand.Intn(80) + 1
	pl(randNum)
	fmt.Printf("%.f", 3.1498)

}
