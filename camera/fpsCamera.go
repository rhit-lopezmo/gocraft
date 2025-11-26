package camera

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type FPSCamera struct {
	Position rl.Vector3
	Yaw      float32
	Pitch    float32

	Speed    float32
	Sens     float32
	MaxPitch float32
}

func NewFPSCamera(pos rl.Vector3) *FPSCamera {
	return &FPSCamera{
		Position: pos,
		Yaw:      0,
		Pitch:    0,
		Speed:    8.0,
		Sens:     0.0025,
		MaxPitch: 1.55, // ~89 degrees
	}
}

func (c *FPSCamera) RaylibCamera() rl.Camera3D {
	return rl.Camera3D{
		Position: c.Position,
		Target: rl.Vector3{
			X: c.Position.X + float32(math.Cos(float64(c.Pitch)))*float32(math.Sin(float64(c.Yaw))),
			Y: c.Position.Y + float32(math.Sin(float64(c.Pitch))),
			Z: c.Position.Z + float32(math.Cos(float64(c.Pitch)))*float32(math.Cos(float64(c.Yaw))),
		},
		Up:         rl.NewVector3(0, 1, 0),
		Fovy:       60.0,
		Projection: rl.CameraPerspective,
	}
}

func (c *FPSCamera) Forward() rl.Vector3 {
	return rl.Vector3{
		X: float32(math.Cos(float64(c.Pitch))) * float32(math.Sin(float64(c.Yaw))),
		Y: float32(math.Sin(float64(c.Pitch))),
		Z: float32(math.Cos(float64(c.Pitch))) * float32(math.Cos(float64(c.Yaw))),
	}
}

func (c *FPSCamera) Update(dt float32) {
	// --------------------------------------------------
	// Mouse Look — correct non-inverted yaw direction
	// --------------------------------------------------
	dx := rl.GetMouseDelta().X
	dy := rl.GetMouseDelta().Y

	c.Yaw -= dx * c.Sens
	c.Pitch -= dy * c.Sens

	if c.Pitch > c.MaxPitch {
		c.Pitch = c.MaxPitch
	}
	if c.Pitch < -c.MaxPitch {
		c.Pitch = -c.MaxPitch
	}

	// --------------------------------------------------
	// Minecraft-style Movement (XZ only) with normalization
	// --------------------------------------------------
	yaw := c.Yaw

	forward := rl.Vector3{
		X: float32(math.Sin(float64(yaw))),
		Y: 0,
		Z: float32(math.Cos(float64(yaw))),
	}

	right := rl.Vector3{
		X: float32(-math.Cos(float64(yaw))),
		Y: 0,
		Z: float32(math.Sin(float64(yaw))),
	}

	move := rl.Vector3{X: 0, Y: 0, Z: 0}

	// Accumulate movement direction
	if rl.IsKeyDown(rl.KeyW) {
		move = rl.Vector3Add(move, forward)
	}
	if rl.IsKeyDown(rl.KeyS) {
		move = rl.Vector3Subtract(move, forward)
	}
	if rl.IsKeyDown(rl.KeyA) {
		move = rl.Vector3Subtract(move, right)
	}
	if rl.IsKeyDown(rl.KeyD) {
		move = rl.Vector3Add(move, right)
	}

	// Vertical movement (optional creative mode)
	if rl.IsKeyDown(rl.KeySpace) {
		move.Y += 1
	}
	if rl.IsKeyDown(rl.KeyLeftShift) {
		move.Y -= 1
	}

	// Normalize diagonal movement
	length := float32(math.Sqrt(float64(move.X*move.X + move.Y*move.Y + move.Z*move.Z)))
	if length > 0.0001 {
		move.X /= length
		move.Y /= length
		move.Z /= length
	}

	// Apply movement
	speed := c.Speed * dt
	c.Position.X += move.X * speed
	c.Position.Y += move.Y * speed
	c.Position.Z += move.Z * speed
}
