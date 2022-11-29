//Евклидово расстояние между двумя точками на плоскости
package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2 float64

	fmt.Scan(&x1, &x2, &y1, &y2)

	evk(x1, x2, y1, y2)

	// выводим результат в os.Stdout
	fmt.Println(evk(x1, x2, y1, y2))
}

func evk(x1, x2, y1, y2 float64) float64 {
	r1 := x1 - x2
	r2 := y1 - y2

	x := math.Pow(r1, 2)
	y := math.Pow(r2, 2)
	res := x + y

	d := math.Sqrt(res)
	return d
}
