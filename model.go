package main

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TickMsg time.Time

type Model struct {
	battleground [][]string
	width        int
	height       int
	borderSymbol string
	invaders     []Invader
	score        int
	gameOver     bool
}

func (m Model) tick() tea.Cmd {
	return tea.Tick(time.Second/3, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return m.tick()

}

func initialModel() Model {
	return Model{
		battleground: [][]string{},
		width:        30,
		height:       20,
		borderSymbol: "#",
		invaders:     []Invader{},
		score:        0,
		gameOver:     false,
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyRunes:
			//attack the invader
			if string(msg.Runes) == m.invaders[len(m.invaders)-1].appearance {
				m.invaders = m.invaders[:len(m.invaders)-1]
				m.score++
			}
			return m, nil
		}
		switch msg.String() {
		//quit game
		case "ctrl+c":
			return m, tea.Quit
		}
	case TickMsg:
		i := NewInvader(m.width)
		m.invaders = append([]Invader{i}, m.invaders...)
		if DetectCollision(&m) {
			m.gameOver = true
			return m, tea.Quit
		}
		return m, m.tick()
	}

	return m, nil
}

func (m Model) View() string {
	var sb strings.Builder

	sb.WriteString(RenderTitle())
	sb.WriteRune('\n')
	var sPlayground strings.Builder
	RenderPlayground(&m)
	RenderInvader(&m)
	for _, row := range m.battleground {
		sPlayground.WriteString(strings.Join(row, ""))
		sPlayground.WriteRune('\n')
	}
	sb.WriteString(sPlayground.String())
	sPlayground.WriteRune('\n')
	sb.WriteString(RenderScore(m.score))
	sb.WriteRune('\n')
	sb.WriteRune('\n')
	sb.WriteString(RenderQuitcommand())
	sb.WriteRune('\n')
	if m.gameOver {
		sb.WriteString(RenderGameOver())
		sb.WriteRune('\n')
	}
	// Send the UI for rendering
	return sb.String()
}
