package cliphandler

import (
	"context"
	"fmt"
	"strings"

	"golang.design/x/clipboard"
)

type Manager struct {
	history []string
	seen    map[string]bool
}

func NewManager() (*Manager, error) {
	if err := clipboard.Init(); err != nil {
		return nil, err
	}

	return &Manager{
		history: make([]string, 0),
		seen:    make(map[string]bool),
	}, nil
}

func (m *Manager) Add(data []byte) {
	if len(data) == 0 {
		return
	}

	text := string(data)
	if m.seen[cleanText(text)] {
		return
	}

	m.seen[cleanText(text)] = true
	m.history = append(m.history, text)

	m.printHistory()
}

func (m *Manager) StartWatching(ctx context.Context) {
	m.Add(clipboard.Read(clipboard.FmtText))

	for data := range clipboard.Watch(ctx, clipboard.FmtText) {
		m.Add(data)
	}
}

func (m *Manager) printHistory() {
	fmt.Println("clipboard history:")
	for i, item := range m.history {
		fmt.Printf("%d: %s\n", i+1, item)
	}
	fmt.Println()
}

func cleanText(sourceText string) string {
	return strings.TrimSpace(strings.ToLower(sourceText))
}
