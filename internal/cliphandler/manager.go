package cliphandler

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"fyne.io/fyne/v2/data/binding"
	"golang.design/x/clipboard"
)

type Manager struct {
	mu       sync.Mutex
	history  []string
	seen     map[string]bool
	DataList binding.StringList
}

func NewManager() (*Manager, error) {
	if err := clipboard.Init(); err != nil {
		return nil, err
	}

	return &Manager{
		history:  make([]string, 0),
		seen:     make(map[string]bool),
		DataList: binding.NewStringList(),
	}, nil
}

func (m *Manager) Add(data []byte) {
	if len(data) == 0 {
		return
	}

	text := string(data)
	cleaned := cleanText(text)

	m.mu.Lock()
	defer m.mu.Unlock()

	if m.seen[cleaned] {
		return
	}

	trimmedText := strings.TrimSpace(text)
	m.seen[cleaned] = true
	m.history = append(m.history, trimmedText)

	_ = m.DataList.Append(trimmedText)

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
