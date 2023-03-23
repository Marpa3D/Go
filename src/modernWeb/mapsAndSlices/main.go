// Хэш-таблицы(карты) и срезы
package main

import (
	"log"
	"sort"
	"strings"
)

type User struct {
	Name     string
	LastName string
}

func main() {
	// Карты
	myMap := make(map[string]string)
	log.Printf("Adress map: %p, value:%+v", &myMap, myMap)

	// Добавление пары ключ:значение
	myMap["cat"] = "Tiger"
	log.Println(myMap["cat"])
	log.Printf("Adress map after init: %p, value:%+v\n", &myMap, myMap)
	log.Printf("Adress map struct init: %p, value:%#v\n", &User{}, User{})

	log.Println(strings.Repeat("*", 50))

	myMap1 := make(map[string]User)
	m := User{
		Name:     "Ilia",
		LastName: "Prorok",
	}
	log.Printf("Adress map after init: %p, value:%+v\n", &m, m)

	log.Println(strings.Repeat("*", 50))

	myMap1["one"] = m
	log.Printf("Adress map after init: %p, value:%+v\n", &myMap1, myMap1["one"].Name)

	log.Println(strings.Repeat("*", 50))

	// Слайсы
	var mySlice []string
	log.Println(mySlice)

	// append
	mySlice = append(mySlice, "Mark")
	log.Printf("Adress: %p, value: %v", &mySlice, mySlice)
	mySlice = append(mySlice, "Shvarchold")
	log.Printf("Adress: %p, value: %v\n", &mySlice, mySlice)

	// Сортировка
	inSlice := []int{4, 3, 2, 5, 1, 7, 8}
	log.Println(inSlice)
	sort.Ints(inSlice)
	log.Println(inSlice)
	log.Println(inSlice[0:2])
	log.Println(inSlice[4:7])

}
