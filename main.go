package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
)

type invader struct {
	apperance string
	x         int
}

func generateInvader() string {
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	invader := rand.Intn(len(letterRunes))

	return string(letterRunes[invader])

}

func newInvader(width int) invader {
	x := rand.Intn(width)

	return invader{
		apperance: generateInvader(),
		x:         x,
	}

}

func detectCollision(m *model) bool {
	return len(m.invaders) == m.height
}

type model struct {
	stopwatch    stopwatch.Model
	playground   [][]string
	width        int
	height       int
	borderSymbol string
	invaders     []invader
	score        int
	gameOver     bool
}

func (m model) Init() tea.Cmd {
	return m.stopwatch.Init()
}

func initialModel() model {
	return model{
		stopwatch:    stopwatch.NewWithInterval(time.Duration(400) * time.Millisecond),
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
		{
			switch msg.String() {

			case "ctrl+c":
				return m, tea.Quit
			}
		}

	}

	var cmd tea.Cmd
	m.stopwatch, cmd = m.stopwatch.Update(msg)
	i := newInvader(m.width)

	m.invaders = append([]invader{i}, m.invaders...)

	if detectCollision(&m) {
		m.gameOver = true
		return m, tea.Quit
	}

	return m, cmd
}

func (m model) View() string {
	// The header
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
