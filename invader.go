package main

import "math/rand"

type Invader struct {
	apperance string
	position  int
}

func GenerateInvader() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	invader := rand.Intn(len(letterRunes))

	return string(letterRunes[invader])

}

func NewInvader(width int) Invader {
	x := rand.Intn(width)

	return Invader{
		apperance: GenerateInvader(),
		position:  x,
	}

}
