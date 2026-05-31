package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func Run(items binding.StringList) {
	myApp := app.New()
	myWindow := myApp.NewWindow("goclip")

	list := widget.NewListWithData(
		items,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(item binding.DataItem, obj fyne.CanvasObject) {
			textBinding := item.(binding.String)
			val, err := textBinding.Get()
			if err != nil {
				return
			}

			obj.(*widget.Label).SetText(val)
		},
	)

	myWindow.SetContent(list)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
