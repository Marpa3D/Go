// Оконное проложение на Go Fyne
// Стартовый шаблон
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"time"
)
func updTime(clock *widget.Label){
	formatTime := time.Now().Format("Текущее время: 03:04:05")
	clock.SetText(formatTime)
}

func main(){
	a := app.New()
	w := a.NewWindow("Крутые часы")

	clock := widget.NewLabel(" ")
	updTime(clock)

	w.SetContent(clock)

	go func(){
		for range time.Tick(time.Second){
			updTime(clock)
		}
	}()
	w.ShowAndRun()

}
