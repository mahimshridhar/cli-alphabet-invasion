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

type model struct {
	stopwatch    stopwatch.Model
	playground   [][]string
	width        int
	height       int
	borderSymbol string
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	// return m.stopwatch.Init()

	//m.currenTetromino = genTetromino()

	return nil
}

func initialModel() model {
	return model{
		stopwatch:    stopwatch.NewWithInterval(time.Duration(80) * time.Millisecond),
		playground:   [][]string{},
		width:        30,
		height:       20,
		borderSymbol: "#",
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		{
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			}
		}

	}
	return m, nil
}

func (m model) View() string {
	// The header
	s := "Welcome to cli tetris\n\n"

	sPlayground := ""

	RenderPlayground(&m)

	for _, row := range m.playground {
		sPlayground += strings.Join(row, "") + "\n"
	}

	s = s + sPlayground

	s = s + "Press q to quit\n\n"

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
