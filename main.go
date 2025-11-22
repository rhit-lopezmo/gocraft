package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "raylib-go window")
	defer rl.CloseWindow()

	// Main game loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Hello, raylib-go!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
