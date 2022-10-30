// Приложение "Время, для приготовления вареного яица" с использованием бибилиотеки Gioui
package main

import (
	"gioui.org/app"
)

func main() {
	go func() {
		w := app.NewWindow()

		for range w.Events() {

		}
	}()
	app.Main()
}
