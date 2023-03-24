package main

import "log"

type HashMap struct {
	buckets [][][]string
	size    int
}

type User struct {
	Name string
	Age  int
}

func main() {
	mySlice := []string{"cat", "dog", "fish", "banana"}
	for i, v := range mySlice {
		log.Println(i, v)
	}

	h := HashMap{
		size: 10,
	}
	log.Printf("Структура бакетов: %#v\n", h)

	var mySlice1 []User

	u1 := User{
		Name: "Mark",
		Age:  44,
	}

	u2 := User{
		Name: "Ilia",
		Age:  44,
	}
	mySlice1 = append(mySlice1, u1, u2)

	for k, v := range mySlice1 {
		log.Println(k, v.Name)
	}
	log.Println(len(mySlice1), cap(mySlice1))
}
