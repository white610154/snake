package main

import (
	"image/color"

	ebi "github.com/hajimehoshi/ebiten"
)

const (
	dirUp = iota
	dirDown
	dirLeft
	dirRight
)

const (
	startSpeed   = 40
	fullDuration = 600
)

// Snake snake that player played
type Snake struct {
	Length    int
	Body      []Point
	Direction int
	Speed     int
	Duration  int
}

// NewSnake create a new snake when game start
func NewSnake() *Snake {
	body := []Point{
		{fieldWidth / 2, fieldHeight / 2},
		{fieldWidth / 2, fieldHeight/2 + 1},
		{fieldWidth / 2, fieldHeight/2 + 2},
		{fieldWidth / 2, fieldHeight/2 + 3},
		{fieldWidth / 2, fieldHeight/2 + 4},
	}
	s := &Snake{5, body, dirUp, startSpeed, fullDuration}
	return s
}

// Image get snake image
func (s *Snake) Image() *ebi.Image {
	node, _ := NewBlock(blockSize, color.RGBA{0, 100, 0, 100})
	img, _ := ebi.NewImage(fieldWidth*blockSize, fieldHeight*blockSize, ebi.FilterDefault)
	for _, p := range s.Body {
		DrawBlock(img, node, p.X, p.Y)
	}
	return img
}

// Move apply snake move
func (s *Snake) Move() bool {
	if s.Duration <= s.Speed {
		head := s.Body[0]
		switch s.Direction {
		case dirUp:
			head.Y--
		case dirDown:
			head.Y++
		case dirLeft:
			head.X--
		case dirRight:
			head.X++
		}
		// if hit the wall
		if head.X < 1 || head.X > fieldWidth || head.Y < 1 || head.Y > fieldHeight {
			return false
		}
		// if hit self
		for _, p := range s.Body {
			if IsEqual(head, p) {
				return false
			}
		}
		// if eat food
		if IsEqual(head, food.Position) {
			food.Generate()
			s.Speed++
		} else {
			s.Body = s.Body[:len(s.Body)-1]
		}
		s.Body = append([]Point{head}, s.Body...)
	}
	s.Duration -= s.Speed
	if s.Duration <= 0 {
		s.Duration += fullDuration
	}
	return true
}

// Eat apply snake eat
func (s *Snake) Eat(dir int) {}
