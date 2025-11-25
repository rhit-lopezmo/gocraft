package camera

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type FPSCamera struct {
	Position    rl.Vector3
	Yaw         float32
	Pitch       float32
	MoveSpeed   float32
	Sensitivity float32
}

func NewFPSCamera(pos rl.Vector3) *FPSCamera {
	return &FPSCamera{
		Position:    pos,
		Yaw:         0,
		Pitch:       0,
		MoveSpeed:   5,
		Sensitivity: 0.003,
	}
}

func (c *FPSCamera) Forward() rl.Vector3 {
	return rl.NewVector3(
		float32(math.Cos(float64(c.Pitch)))*float32(math.Sin(float64(c.Yaw))),
		float32(math.Sin(float64(c.Pitch))),
		float32(math.Cos(float64(c.Pitch)))*float32(math.Cos(float64(c.Yaw))),
	)
}

func (c *FPSCamera) Right() rl.Vector3 {
	f := c.Forward()
	return rl.NewVector3(-f.Z, 0, f.X)
}

func (c *FPSCamera) Up() rl.Vector3 {
	return rl.NewVector3(0, 1, 0)
}

func (c *FPSCamera) Update(dt float32) {
	mouse := rl.GetMouseDelta()

	c.Yaw -= mouse.X * c.Sensitivity
	c.Pitch -= mouse.Y * c.Sensitivity

	limit := float32(math.Pi/2 - 0.01)
	if c.Pitch > limit {
		c.Pitch = limit
	}
	if c.Pitch < -limit {
		c.Pitch = -limit
	}

	if c.Yaw > math.Pi {
		c.Yaw -= 2 * math.Pi
	} else if c.Yaw < -math.Pi {
		c.Yaw += 2 * math.Pi
	}

	move := rl.NewVector3(0, 0, 0)
	f := c.Forward()
	r := c.Right()

	if rl.IsKeyDown(rl.KeyW) {
		move = rl.Vector3Add(move, f)
	}
	if rl.IsKeyDown(rl.KeyS) {
		move = rl.Vector3Subtract(move, f)
	}
	if rl.IsKeyDown(rl.KeyD) {
		move = rl.Vector3Add(move, r)
	}
	if rl.IsKeyDown(rl.KeyA) {
		move = rl.Vector3Subtract(move, r)
	}

	if rl.Vector3Length(move) > 0 {
		move = rl.Vector3Normalize(move)
		move = rl.Vector3Scale(move, c.MoveSpeed*dt)
		c.Position = rl.Vector3Add(c.Position, move)
	}
}

func (c *FPSCamera) RaylibCamera() rl.Camera3D {
	f := c.Forward()
	return rl.Camera3D{
		Position:   c.Position,
		Target:     rl.Vector3Add(c.Position, f),
		Up:         rl.NewVector3(0, 1, 0),
		Fovy:       75,
		Projection: rl.CameraPerspective,
	}
}
