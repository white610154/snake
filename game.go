package main

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/ebitenutil"

	ebi "github.com/hajimehoshi/ebiten"
)

const (
	windowWidth    = 1920
	windowHeight   = 1080
	fieldWidth     = 30
	fieldHeight    = 30
	blockSize      = 20
	fieldLeftShift = (windowWidth - fieldWidth*blockSize) / 2
	fieldUpShift   = (windowHeight - fieldHeight*blockSize) / 2
)

var (
	snake         *Snake
	food          *Food
	err           error
	field         *ebi.Image
	fieldPosition *ebi.DrawImageOptions

	running = false
	msg     string

	print = ebitenutil.DebugPrint
)

// Game game object for ebiten
type Game struct{}

// Update the logical state
func (g *Game) Update(screen *ebi.Image) error {
	if ebi.IsKeyPressed(ebi.KeyEscape) {
		os.Exit(0)
	}
	if ebi.IsKeyPressed(ebi.KeyEnter) {
		snake = NewSnake()
		food = NewFood()
		if err != nil {
			panic(err)
		}
		running = true
	}
	if running {
		if ebi.IsKeyPressed(ebi.KeyW) && snake.Direction != dirDown {
			snake.Direction = dirUp
		}
		if ebi.IsKeyPressed(ebi.KeyS) && snake.Direction != dirUp {
			snake.Direction = dirDown
		}
		if ebi.IsKeyPressed(ebi.KeyA) && snake.Direction != dirRight {
			snake.Direction = dirLeft
		}
		if ebi.IsKeyPressed(ebi.KeyD) && snake.Direction != dirLeft {
			snake.Direction = dirRight
		}
		move := snake.Move()
		if !move {
			msg = "Game Over!! Press 'Enter' to play again."
			running = false
		}
	}
	return nil
}

// Draw render the screen
func (g *Game) Draw(screen *ebi.Image) {
	screen.Fill(color.Black)
	if running {
		screen.DrawImage(field, nil)
		screen.DrawImage(snake.Image(), fieldPosition)
		screen.DrawImage(food.Image, food.Option)
	}
	print(screen, msg)
}

// Layout return the screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func main() {
	fieldPosition = &ebi.DrawImageOptions{}
	fieldPosition.GeoM.Translate(fieldLeftShift, fieldUpShift)
	msg = "Press 'Enter' to Start"

	ebi.SetFullscreen(true)
	ebi.SetWindowSize(1920, 1080)
	ebi.SetWindowTitle("貪食蛇")
	game := &Game{}
	if err := ebi.RunGame(game); err != nil {
		panic(err)
	}
}
