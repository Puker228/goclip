package ui

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
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

func Run(items binding.StringList, onCopy func(string), onClean func()) {
	myApp := app.New()
	myWindow := myApp.NewWindow("goclip")

	list := widget.NewListWithData(
		items,
		func() fyne.CanvasObject {
			label := widget.NewLabel("template")
			label.Wrapping = fyne.TextWrapOff
			label.Truncation = fyne.TextTruncateEllipsis

			copyButton := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), nil)
			copyButton.Importance = widget.LowImportance

			return container.NewBorder(nil, nil, nil, copyButton, label)
		},
		func(item binding.DataItem, obj fyne.CanvasObject) {
			textBinding := item.(binding.String)
			val, err := textBinding.Get()
			if err != nil {
				return
			}

			row := obj.(*fyne.Container)
			label := row.Objects[0].(*widget.Label)
			copyButton := row.Objects[1].(*widget.Button)

			label.SetText(preview(val, 100))
			copyButton.OnTapped = func() {
				onCopy(val)
			}
		},
	)

	cleanButton := widget.NewButtonWithIcon("Clean", theme.DeleteIcon(), onClean)
	cleanButton.Importance = widget.HighImportance

	footer := container.NewBorder(nil, nil, nil, cleanButton)

	myWindow.SetContent(container.NewBorder(nil, footer, nil, nil, list))
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
