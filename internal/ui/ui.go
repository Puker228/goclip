package ui

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func preview(s string, max int) string {
	s = strings.Join(strings.Fields(s), " ")

	r := []rune(s)
	if len(r) <= max {
		return s
	}
	return string(r[:max]) + "…"
}

func Run(items binding.StringList) {
	myApp := app.New()
	myWindow := myApp.NewWindow("goclip")

	list := widget.NewListWithData(
		items,
		func() fyne.CanvasObject {
			label := widget.NewLabel("template")
			label.Wrapping = fyne.TextWrapOff
			label.Truncation = fyne.TextTruncateEllipsis

			return label
		},
		func(item binding.DataItem, obj fyne.CanvasObject) {
			textBinding := item.(binding.String)
			val, err := textBinding.Get()
			if err != nil {
				return
			}

			obj.(*widget.Label).SetText(preview(val, 100))
		},
	)

	myWindow.SetContent(list)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
