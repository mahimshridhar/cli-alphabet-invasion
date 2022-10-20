package main

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TickMsg time.Time

func detectCollision(m *Model) bool {
	return len(m.invaders) == m.height
}

type Model struct {
	batteground  [][]string
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
		batteground:  [][]string{},
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
			if string(msg.Runes) == m.invaders[len(m.invaders)-1].apperance {
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

		if detectCollision(&m) {
			m.gameOver = true
			return m, tea.Quit

		}

		return m, m.tick()

	}

	return m, nil
}

func (m Model) View() string {
	// var sb strings.Builder

	s := RenderTitle() + "\n"

	sPlayground := ""

	RenderPlayground(&m)

	RenderInvader(&m)

	for _, row := range m.batteground {
		sPlayground += strings.Join(row, "") + "\n"
	}

	s = s + sPlayground

	s = s + RenderScore(m.score) + "\n"

	s = s + RenderQuitcommand() + "\n"

	if m.gameOver {
		s = s + RenderGameOver() + "\n"
	}

	// Send the UI for rendering
	return s
}
