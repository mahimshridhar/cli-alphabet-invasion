package main

func DetectCollision(m *Model) bool {
	return len(m.invaders) == m.height
}
