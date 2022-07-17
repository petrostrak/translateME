package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	win := a.NewWindow("translateMe")

	from, to, input, output := makeUI()

	win.SetContent(container.NewHBox(from, to, container.NewVSplit(input, output)))

	win.Resize(fyne.Size{Width: 800, Height: 500})
	win.CenterOnScreen()
	win.ShowAndRun()
}
