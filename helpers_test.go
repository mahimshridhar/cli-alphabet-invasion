package main

import "testing"

func TestCollisionDetection(t *testing.T) {
	h := 10
	var invaders []Invader
	for i := 0; i < h; i++ {
		invaders = append(invaders, Invader{})
	}
	m := Model{
		height:   h,
		invaders: invaders,
	}

	if DetectCollision(&m) == false {
		t.Error("Collision detection not working")
	}
}
