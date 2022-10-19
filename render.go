package main

import "strings"

func RenderPlayground(m *model) {
	for i := 0; i < m.height; i++ {
		m.playground = append(m.playground, strings.Split(strings.Repeat(" ", m.width), ""))
	}

	m.playground = append(m.playground, strings.Split(strings.Repeat(m.borderSymbol, m.width), ""))

}
