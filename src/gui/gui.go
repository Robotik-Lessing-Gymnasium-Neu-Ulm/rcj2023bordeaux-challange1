package gui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func InitGui() {
	a := app.New()
	w := a.NewWindow("Robot Controller")

	empty := widget.NewLabel(" ")
	wt := widget.NewLabel("Robot Controller - RCJ2023 Bordeaux GPS Challange")

	p1 := widget.NewButton("1", func() {
	})

	p2 := widget.NewButton("2", func() {
	})

	p3 := widget.NewButton("3", func() {
	})

	p4 := widget.NewButton("4", func() {
	})

	p5 := widget.NewButton("5", func() {
	})

	grid := container.New(layout.NewGridLayout(5), empty, empty, wt, empty, empty, p1, p2, p3, p4, p5)

	w.SetContent(grid)
	w.ShowAndRun()
}
