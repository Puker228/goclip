package cliphandler

import (
	"context"
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

func (m *Manager) Copy(text string) {
	clipboard.Write(clipboard.FmtText, []byte(text))
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
}

func (m *Manager) Clean() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.history = m.history[:0]
	m.seen = make(map[string]bool)
	_ = m.DataList.Set([]string{})
}

func (m *Manager) StartWatching(ctx context.Context) {
	m.Add(clipboard.Read(clipboard.FmtText))

	for data := range clipboard.Watch(ctx, clipboard.FmtText) {
		m.Add(data)
	}
}

func cleanText(sourceText string) string {
	return strings.TrimSpace(strings.ToLower(sourceText))
}
