package main

import (
	"context"
	"fmt"

	"golang.design/x/clipboard"
)

func main() {
	if err := clipboard.Init(); err != nil {
		panic(err)
	}

	var history []string
	seen := map[string]bool{}

	add := func(data []byte) {
		if len(data) == 0 {
			return
		}

		text := string(data)
		if seen[text] {
			return
		}

		seen[text] = true
		history = append(history, text)

		fmt.Println("clipboard history:")
		for i, item := range history {
			fmt.Printf("%d: %s\n", i+1, item)
		}
		fmt.Println()
	}

	add(clipboard.Read(clipboard.FmtText))

	for data := range clipboard.Watch(context.Background(), clipboard.FmtText) {
		add(data)
	}
}
