// Массивы. Курс от "Специалист"
package main

import "fmt"

func main() {
	// Массив - это статическая последовательность элементов в памяти с фиксированным колличеством этих элементов
	var arr [8]int // При инициализации массива важно указывать сколько будет в нем элементов! индекс 0 - первый элемент,
	// последний элемент - это len - 1
	fmt.Println("Arrays:", arr)
	arr[1] = 720
	fmt.Println("Arrays:", arr)

	newArr := [5]float64{12.0, 13.4, 10.0, 22.8, 88.8}
	fmt.Println("Arrays:", newArr)

	// Создание массива через инициализацию переменных

	arrWithValues := [...]int{}
	fmt.Println("Arrays:", arrWithValues, "Длинна:", len(arrWithValues))

	// Массив - это набор ЗНАЧЕНИЙ!!! Пи всех манипуляциях массив копируется жестко!
	first := [...]int{1, 2, 3, 4}
	second := [...]int{5, 6, 7, 8}
	first = second
	second[1] = 720
	fmt.Println("Первый массив:", first, "Второй массив:", second)
	// Массив и его размер - это две составляющие его типа.Размер массива - это часть его типа!
	// arrFirst [4]int != arr arrSecond [5]int (это разные типы!)

	// Итерирование по массиву
	newArr = [5]float64{12.0, 13.4, 10.0, 22.8, 88.8}
	for i := 0; i < len(newArr); i++ {
		fmt.Printf("%d elements of arr is % 0.2f\n", i, newArr[i])
	}
	var sum float64
	for j, value := range newArr {
		fmt.Printf("%d element or %0.2f\n", j, value)
		sum += value
		fmt.Println(sum)
	}

	// Игнорирование id in range

	for _, val := range newArr {
		fmt.Printf("%0.2f value\n", val)
	}

	// Многомерные массивы. Должны содержать в себе элементы одного типа!

	words := [2][2]string{
		{"John", "Lisa"},
		{"Mark", "Matwey"},
	}
	fmt.Println(words)

}
