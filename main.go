package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

const (
	mouseSensitivity = 0.003
	moveSpeed        = 5.0
)

var (
	yaw   float32
	pitch float32
)

func main() {
	screenWidth := int32(1280)
	screenHeight := int32(720)

	rl.InitWindow(screenWidth, screenHeight, "Gocraft")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	rl.DisableCursor()

	camera := rl.Camera3D{
		Position:   rl.NewVector3(0.0, 1.8, 0.0),
		Target:     rl.NewVector3(0.0, 1.8, 1.0),
		Up:         rl.NewVector3(0.0, 1.0, 0.0),
		Fovy:       75.0,
		Projection: rl.CameraPerspective,
	}

	for !rl.WindowShouldClose() {
		update(&camera)
		draw(camera)
	}
}

func update(camera *rl.Camera3D) {
	dt := rl.GetFrameTime()

	// === Mouse Look ===
	mouseDelta := rl.GetMouseDelta()

	yaw -= mouseDelta.X * mouseSensitivity
	pitch -= mouseDelta.Y * mouseSensitivity

	// Clamp pitch
	maxPitch := float32(math.Pi/2 - 0.01)
	if pitch > maxPitch {
		pitch = maxPitch
	}
	if pitch < -maxPitch {
		pitch = -maxPitch
	}

	// Direction vectors
	forward := rl.NewVector3(
		float32(math.Cos(float64(pitch)))*float32(math.Sin(float64(yaw))),
		float32(math.Sin(float64(pitch))),
		float32(math.Cos(float64(pitch)))*float32(math.Cos(float64(yaw))),
	)

	right := rl.NewVector3(
		float32(math.Sin(float64(yaw)-math.Pi/2)),
		0,
		float32(math.Cos(float64(yaw)-math.Pi/2)),
	)

	// === Movement ===
	movement := rl.NewVector3(0, 0, 0)

	if rl.IsKeyDown(rl.KeyW) {
		movement = rl.Vector3Add(movement, forward)
	}
	if rl.IsKeyDown(rl.KeyS) {
		movement = rl.Vector3Subtract(movement, forward)
	}
	if rl.IsKeyDown(rl.KeyD) {
		movement = rl.Vector3Add(movement, right)
	}
	if rl.IsKeyDown(rl.KeyA) {
		movement = rl.Vector3Subtract(movement, right)
	}

	// Normalize + scale
	if rl.Vector3Length(movement) > 0 {
		movement = rl.Vector3Normalize(movement)
		movement = rl.Vector3Scale(movement, moveSpeed*dt)
	}

	// Update camera
	camera.Position = rl.Vector3Add(camera.Position, movement)
	camera.Target = rl.Vector3Add(camera.Position, forward)
}

func draw(camera rl.Camera3D) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode3D(camera)

	rl.DrawCube(rl.NewVector3(0, 0.5, 5), 1, 1, 1, rl.Blue)
	rl.DrawGrid(20, 1.0)

	rl.EndMode3D()

	rl.DrawFPS(10, 10)
	rl.EndDrawing()
}
