// Несколько трюков со срезами
package main

import "fmt"

// Обычный способ
//func addUsers(users []string) {
//	for _, user := range users {
//		fmt.Println(user)
//	}
//}
//
//func main() {
//	addUsers([]string{"John", "Mark", "Lary"})
//}

// Трюк 1. Функция с переменным числом аргументов
//func addUsers(users ...string) {
//	for _, user := range users {
//		fmt.Println(user)
//	}
//}
//
//func main() {
//	addUsers("John", "Mark", "Lary")
//}

// Трюк 2.

var users = []string{}

func addUsers(user ...string) {
	users = append(users, user...)
}

func main() {
	addUsers("John", "Mark", "Lary", "Лука")
	fmt.Println(users)
}
