package main

import (
	"image/color"

	ebi "github.com/hajimehoshi/ebiten"
)

// Food food object properties
type Food struct {
	Position Point
	Image    *ebi.Image
	Option   *ebi.DrawImageOptions
}

// NewFood initialize a food object
func NewFood() *Food {
	f := &Food{}
	f.Image, _ = NewBlock(blockSize, color.RGBA{100, 100, 0, 200})
	f.Generate()
	return f
}

// Generate generate drawImageOptions for new food
func (f *Food) Generate() {
	f.Position = RandomPoint(snake.Body...)
	op := &ebi.DrawImageOptions{}
	op.GeoM.Translate(float64((f.Position.X-1)*blockSize+fieldLeftShift), float64((f.Position.Y-1)*blockSize+fieldUpShift))
	f.Option = op
}
