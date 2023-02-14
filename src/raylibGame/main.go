// game raylib with Go
package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - начальное окно")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Welcome!!! This is first game window!", 190, 200, 20, rl.DarkGray)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
