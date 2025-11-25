package main

import (
	"fmt"
	"gocraft/camera"
	"gocraft/terrain"
	"math"
	"unsafe"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	mesh     rl.Mesh
	material rl.Material
)

func main() {
	screenWidth := int32(1280)
	screenHeight := int32(720)

	rl.InitWindow(screenWidth, screenHeight, "Gocraft")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	rl.DisableCursor()

	cam := camera.NewFPSCamera(rl.NewVector3(0, 6, -6))

	chunk := terrain.NewChunk(0, 0)
	chunk.GenerateFlat(4)
	mesh = chunk.GenerateMesh()

	maps := make([]rl.MaterialMap, 12)
	maps[0].Color = rl.NewColor(120, 200, 80, 255)

	material = rl.Material{}
	material.Shader = rl.LoadShader("", "")
	material.Maps = &maps[0]

	defer rl.UnloadMesh(&mesh)
	defer rl.UnloadMaterial(material)

	for !rl.WindowShouldClose() {
		update(cam)
		draw(cam)
	}
}

func update(cam *camera.FPSCamera) {
	dt := rl.GetFrameTime()
	cam.Update(dt)
}

func draw(cam *camera.FPSCamera) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode3D(cam.RaylibCamera())
	rl.DrawMesh(mesh, material, rl.MatrixTranslate(0, 0, 0))
	drawTrianglesDebug(mesh)
	rl.DrawGrid(50, 1)
	rl.EndMode3D()

	rl.DrawFPS(10, 10)
	debugCameraInfo(cam)

	rl.EndDrawing()
}

func drawTrianglesDebug(mesh rl.Mesh) {
	if mesh.Vertices == nil || mesh.Indices == nil {
		return
	}

	// Convert C-style pointer → Go slice using SliceHeader
	vertCount := int(mesh.VertexCount * 3)  // xyz per vertex
	triCount := int(mesh.TriangleCount * 3) // 3 indices per triangle

	vertSlice := unsafe.Slice(mesh.Vertices, vertCount)
	indSlice := unsafe.Slice(mesh.Indices, triCount)

	for i := 0; i < len(indSlice); i += 3 {
		i0 := int(indSlice[i]) * 3
		i1 := int(indSlice[i+1]) * 3
		i2 := int(indSlice[i+2]) * 3

		v0 := rl.NewVector3(vertSlice[i0], vertSlice[i0+1], vertSlice[i0+2])
		v1 := rl.NewVector3(vertSlice[i1], vertSlice[i1+1], vertSlice[i1+2])
		v2 := rl.NewVector3(vertSlice[i2], vertSlice[i2+1], vertSlice[i2+2])

		// Fill triangle (optional, for visibility)
		rl.DrawTriangle3D(v0, v1, v2, rl.NewColor(255, 0, 0, 80))

		// Wireframe outline
		rl.DrawLine3D(v0, v1, rl.Black)
		rl.DrawLine3D(v1, v2, rl.Black)
		rl.DrawLine3D(v2, v0, rl.Black)
	}
}

func debugCameraInfo(cam *camera.FPSCamera) {
	f := cam.Forward()

	yawDeg := cam.Yaw * 180 / float32(math.Pi)
	pitchDeg := cam.Pitch * 180 / float32(math.Pi)

	for yawDeg > 180 {
		yawDeg -= 360
	}
	for yawDeg < -180 {
		yawDeg += 360
	}

	var horiz string
	switch {
	case yawDeg >= -45 && yawDeg < 45:
		horiz = "+Z" // forward
	case yawDeg >= 45 && yawDeg < 135:
		horiz = "-X" // left
	case yawDeg >= -135 && yawDeg < -45:
		horiz = "+X" // right
	default:
		horiz = "-Z" // back
	}
	var vert string
	if pitchDeg > 30 {
		vert = "+Y"
	} else if pitchDeg < -30 {
		vert = "-Y"
	} else {
		vert = "0Y"
	}

	rl.DrawText(fmt.Sprintf("Forward: (%.2f %.2f %.2f)", f.X, f.Y, f.Z), 10, 40, 20, rl.Black)
	rl.DrawText(fmt.Sprintf("Yaw:   %.1f°", yawDeg), 10, 65, 20, rl.Black)
	rl.DrawText(fmt.Sprintf("Pitch: %.1f°", pitchDeg), 10, 90, 20, rl.Black)
	rl.DrawText("Facing: "+horiz+"  "+vert, 10, 115, 20, rl.Black)
}
