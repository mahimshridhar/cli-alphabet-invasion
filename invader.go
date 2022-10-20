package main

import "math/rand"

type Invader struct {
	appearance string
	position   int
}

const (
	ALPHABET_INVADERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GenerateInvader() string {
	invader := rand.Intn(len(ALPHABET_INVADERS))
	return string(ALPHABET_INVADERS[invader])

}

func NewInvader(width int) Invader {
	x := rand.Intn(width)

	return Invader{
		appearance: GenerateInvader(),
		position:   x,
	}

}
