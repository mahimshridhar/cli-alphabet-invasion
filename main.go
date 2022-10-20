package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TickMsg time.Time

type invader struct {
	apperance string
	position  int
}

func generateInvader() string {
	//invaders
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	invader := rand.Intn(len(letterRunes))

	return string(letterRunes[invader])

}

func newInvader(width int) invader {
	x := rand.Intn(width)

	return invader{
		apperance: generateInvader(),
		position:  x,
	}

}

func detectCollision(m *model) bool {
	return len(m.invaders) == m.height
}

type model struct {
	playground   [][]string
	width        int
	height       int
	borderSymbol string
	invaders     []invader
	score        int
	gameOver     bool
}

func (m model) tick() tea.Cmd {
	return tea.Tick(time.Second/3, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m model) Init() tea.Cmd {
	return m.tick()

}

func initialModel() model {
	return model{
		playground:   [][]string{},
		width:        30,
		height:       20,
		borderSymbol: "#",
		invaders:     []invader{},
		score:        0,
		gameOver:     false,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		i := newInvader(m.width)

		m.invaders = append([]invader{i}, m.invaders...)

		if detectCollision(&m) {
			m.gameOver = true
			return m, tea.Quit

		}

		return m, m.tick()

	}

	return m, nil
}

func (m model) View() string {
	s := "Welcome to cli tetris\n\n"

	sPlayground := ""

	RenderPlayground(&m)

	RenderInvader(&m)

	for _, row := range m.playground {
		sPlayground += strings.Join(row, "") + "\n"
	}

	s = s + sPlayground

	s = s + "Press ctrl + c to quit\n\n"

	s = s + RenderScore(m.score)

	if m.gameOver {
		s = s + "\nGame Over.\n\n"
	}

	// Send the UI for rendering
	return s
}

func main() {
	rand.Seed(time.Now().UnixNano())
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
