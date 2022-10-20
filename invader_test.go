package main

import (
	"strings"
	"testing"
)

func TestNewInvaderPosition(t *testing.T) {
	i := NewInvader(50)

	if i.position > 50 {
		t.Error("Invader is not being created within the battefeild limit")
	}
}

func TestNewInvaderApperance(t *testing.T) {
	i := NewInvader(50)

	if !strings.Contains(ALPHABET_INVADERS, i.appearance) {
		t.Error("Invader has a wrong appearance.")
	}
}

func TestInvaderAttacked(t *testing.T) {
	keyPressed := "A"

	m := Model{
		invaders: []Invader{{appearance: "C"}, {appearance: "A"}},
	}

	if keyPressed == m.invaders[len(m.invaders)-1].appearance {
		m.invaders = m.invaders[:len(m.invaders)-1]

	}

	if len(m.invaders) != 1 {
		t.Error("Invader attack has failed.")

	}
}
