// Циклы
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var pl = fmt.Println
var plf = fmt.Printf

func main() {
	for i := 0; i <= 10; i++ {
		plf("%+v\t", i)
	}
	pl("\n")
	for j := 10; j >= 0; j-- {
		plf("%d\t", j)
	}
	pl("\n")
	// while
	var w int
	for w < 10 {
		plf("%d\t", w)
		if w == 8 {
			break
		}
		w++
	}
	pl("\n")

	// Случайное число в циклах
	seekSecs := time.Now().Unix()
	rand.Seed(seekSecs)
	randNum := rand.Intn(100) + 1
	for true {
		fmt.Print("Угадайте число от 0 до 100:")
		pl("Секретное число:", randNum)

		// Ввод от пользователя. Игра "Угадай число"
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		strInt, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if strInt > randNum {
			pl("Ваше число больше секретного.")
		} else if strInt < randNum {
			pl("Ваше число меньше секретного.")
		} else {
			pl("Вы угадали секретное число. С Победой!!!")
			break
		}
	}

	// Итерация. Range
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i, v := range slice {
		plf("%d : %+v : %p\n", i, v, &slice[i]) // печать адресов среза и его елементов в памяти
	}

}
