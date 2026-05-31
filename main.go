package main

import (
	"context"
	"log"

	"github.com/Puker228/goclip/internal/cliphandler"
	"github.com/Puker228/goclip/internal/ui"
)

func main() {
	mgr, err := cliphandler.NewManager()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Println("Service started")

	go mgr.StartWatching(context.Background())

	ui.Run(mgr.DataList)
}
