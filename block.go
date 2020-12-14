package main

import (
	"image/color"

	ebi "github.com/hajimehoshi/ebiten"
)

// NewBlock create a block
func NewBlock(size int, clr color.Color) (*ebi.Image, error) {
	blockImage, err := ebi.NewImage(size, size, ebi.FilterDefault)
	if err != nil {
		return nil, err
	}
	tempImage, err := ebi.NewImage(size-4, size-4, ebi.FilterDefault)
	if err != nil {
		return nil, err
	}
	tempImage.Fill(clr)
	op := &ebi.DrawImageOptions{}
	op.GeoM.Translate(2, 2)
	blockImage.DrawImage(tempImage, op)
	return blockImage, nil
}

// DrawBlock draw a block on the image with translate
func DrawBlock(image *ebi.Image, block *ebi.Image, x, y int) *ebi.Image {
	op := &ebi.DrawImageOptions{}
	op.GeoM.Translate(float64((x-1)*blockSize), float64((y-1)*blockSize))
	image.DrawImage(block, op)
	return image
}
