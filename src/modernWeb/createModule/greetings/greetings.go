package greetings

import (
	"errors"
	"fmt"
)

func Hello(message string) (string, error) {

	if message == "" {
		return "", errors.New("Имя не задано!")
	}
	msg := fmt.Sprintf("Привет, %v. Добро пожаловать!", message)
	return msg, nil
}
