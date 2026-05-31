package main

import (
	"context"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/Puker228/goclip/internal/cliphandler"
)

func main() {
	mgr, err := cliphandler.NewManager()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Println("Service started")

	go mgr.StartWatching(context.Background())

	myApp := app.New()
	myWindow := myApp.NewWindow("goclip")

	list := widget.NewListWithData(
		mgr.DataList,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(item binding.DataItem, obj fyne.CanvasObject) {
			textBinding := item.(binding.String)
			val, err := textBinding.Get()
			if err != nil {
				return
			}
			// Присваиваем текст нашему лейблу
			obj.(*widget.Label).SetText(val)
		},
	)

	myWindow.SetContent(list)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
