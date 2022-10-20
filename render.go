package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func RenderPlayground(m *model) {
	for i := 0; i < m.height; i++ {
		m.playground = append(m.playground, strings.Split(strings.Repeat(" ", m.width), ""))
	}
	m.playground = append(m.playground, strings.Split(strings.Repeat(m.borderSymbol, m.width), ""))

}

func RenderInvader(m *model) {
	for i, val := range m.invaders {
		m.playground[i][val.position] = val.apperance
	}
}

func RenderScore(score int) string {

	scoreStr := fmt.Sprintf("Score: %d ", score)
	ts := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10"))

	return ts.Render(scoreStr)
}

func RenderTitle() string {
	ts := lipgloss.NewStyle().Bold(true).
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("63")).
		Width(30).
		AlignHorizontal(lipgloss.Center).
		MarginTop(1).
		MarginBottom(1).
		Underline(true)
	return ts.Render("ALPHABET INVASION")

}

func RenderQuitcommand() string {
	qc := "Press ctrl+c to quit"
	ts := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63"))
	return ts.Render((qc))
}

func RenderGameOver() string {
	return lipgloss.NewStyle().Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Width(30).
		AlignHorizontal(lipgloss.Center).
		MarginTop(1).
		MarginBottom(1).
		Render("Game Over!")
}
