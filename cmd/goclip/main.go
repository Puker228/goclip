package main

import (
	"context"
	"log"

	"github.com/Puker228/goclip/internal/cliphandler"
)

func main() {
	mgr, err := cliphandler.NewManager()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Println("Service started")

	mgr.StartWatching(context.Background())
}
