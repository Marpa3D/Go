// Оконное проложение на Go Fyne
// Стартовый шаблон
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)
func main(){
	a := app.New()
	w := a.NewWindow("Первое оконное приложение на GO!!!")
	w.SetContent(widget.NewLabel("Приветики!"))
	w.ShowAndRun()

}
