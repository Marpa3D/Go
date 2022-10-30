// Приложение "Время, для приготовления вареного яица" с использованием бибилиотеки Gioui
package main

import (
	"gioui.org/app"
	"gioui.org/unit"
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Title("Таймер готовки"),          // Название приложения
			app.Size(unit.Dp(400), unit.Dp(600)), // размеры окна
		)

		for range w.Events() {

		}
	}()
	app.Main()
}
