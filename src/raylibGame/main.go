// game raylib with Go
package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 1000
	screenHeigth = 480
)

var (
	running      = true
	bgColor      = rl.NewColor(147, 211, 196, 255)
	grassSprite  rl.Texture2D
	playerSprite rl.Texture2D
)

func init() {
	rl.InitWindow(screenWidth, screenHeigth, "Стартовый экран")
	rl.SetExitKey(0) // окно не закроется по нажатию [esc]. Только мышью на [X]
	grassSprite = rl.LoadTexture("resurses/Tilesets/groundTiles/oldTiles/Grass.png")
	playerSprite = rl.LoadTexture("resurses/Characters/CharacterSpriteSheet.png")
	rl.SetTargetFPS(60)
}

func drawScene() {
	rl.DrawTexture(grassSprite, 100, 100, rl.White)
	rl.DrawTexture(playerSprite, 200, 100, rl.White)
}

func input() {}

func update() {
	running = !rl.WindowShouldClose()
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bgColor)
	drawScene()
	rl.EndDrawing()
}

func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.CloseWindow()
}

func main() {

	for running {
		input()
		update()
		render()
	}

	quit()
}
