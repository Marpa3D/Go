// Приложение "Время, для приготовления вареного яица" с использованием бибилиотеки Gioui
package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"os"
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Title("Таймер готовки"),          // Название приложения
			app.Size(unit.Dp(400), unit.Dp(600)), // размеры окна
		)
		// Перемнная ops определяет операции из пользовательского интерфейса
		var ops op.Ops

		// кнопка - кликабельный виджет
		var startButton widget.Clickable

		// th - переменная. Опроеделяет стиль Material Design (тему) и устанавливает шрифт, как gofonts
		th := material.NewTheme(gofont.Collection())

		// Бесконечный цикл. Слушаем события в окне приложения. w.Events - это канал, по нему добавляются события
		for e := range w.Events() {
			// e.(type) - это переключатель типов. Позволяет выполнять различные действия, в зависимости от (type) события
			switch e := e.(type) {
			// Это сообщение отправляется, когда приложение должно выполнить повторный рендеринг
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				btn := material.Button(th, &startButton, "Старт")
				btn.Layout(gtx)
				e.Frame(gtx.Ops)
			}
		}
		os.Exit(0)
	}()
	app.Main()
}
