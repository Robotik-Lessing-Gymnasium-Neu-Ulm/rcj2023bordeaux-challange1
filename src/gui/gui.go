package gui

import (
	"rcj2023bordeau-challange1/src/teensy"

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
		teensy.CalculateDirection(1)
	})

	p2 := widget.NewButton("2", func() {
		teensy.CalculateDirection(2)

	})

	p3 := widget.NewButton("3", func() {
		teensy.CalculateDirection(3)

	})

	p4 := widget.NewButton("4", func() {
		teensy.CalculateDirection(4)

	})

	p5 := widget.NewButton("5", func() {
		teensy.CalculateDirection(5)

	})

	grid := container.New(layout.NewGridLayout(5), empty, empty, wt, empty, empty, p1, p2, p3, p4, p5)

	w.SetContent(grid)
	w.ShowAndRun()
}
