package greetings

import "fmt"

func Hello(message string) string {
	msg := fmt.Sprintf("Привет, %v. Добро пожаловать!", message)
	return msg
}
