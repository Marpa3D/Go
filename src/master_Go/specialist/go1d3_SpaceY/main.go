package main

import (
	"fmt"
)

func main() {
	var n, min, max float64
	//var count int
	const minDia = 100.0
	const maxDia = 140.0

	var sl []float64

	for {
		fmt.Scan(&n)

		if n >= minDia && n <= maxDia {
			sl = append(sl, n)
		}

		//min = sl[0]
		//max = sl[0]

		if n < 0 {
			min = sl[0]
			max = sl[0]
			for _, v := range sl {
				if v < min {
					min = v
				}
				if v > max {
					max = v
				}
			}
			//fmt.Println(count)
			fmt.Println(len(sl))
			fmt.Println(min, max)
			break
		}

	}
}
