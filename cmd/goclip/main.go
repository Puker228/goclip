package main

import (
	"fmt"

	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	data := clipboard.Read(clipboard.FmtText)

	if len(data) == 0 {
		fmt.Println("clipboard is empty")
	} else {
		fmt.Printf("last clipboard element: %s\n", string(data))
	}
}
