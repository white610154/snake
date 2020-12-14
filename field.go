package main

import (
	"image/color"

	ebi "github.com/hajimehoshi/ebiten"
)

func init() {
	field, _ = ebi.NewImage(windowWidth, windowHeight, ebi.FilterDefault)
	wall, _ := NewBlock(blockSize, color.White)
	realField, _ := ebi.NewImage((fieldWidth+2)*blockSize, (fieldHeight+2)*blockSize, ebi.FilterDefault)
	for i := 1; i <= fieldWidth+2; i++ {
		DrawBlock(realField, wall, i, 1)
		DrawBlock(realField, wall, i, fieldHeight+2)
	}
	for i := 1; i <= fieldHeight; i++ {
		DrawBlock(realField, wall, 1, i+1)
		DrawBlock(realField, wall, fieldWidth+2, i+1)
	}
	op := &ebi.DrawImageOptions{}
	op.GeoM.Translate(fieldLeftShift-blockSize, fieldUpShift-blockSize)
	field.DrawImage(realField, op)
}
