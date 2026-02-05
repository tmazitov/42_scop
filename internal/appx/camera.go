package appx

import (
    "github.com/go-gl/mathgl/mgl32"
	"math"
)

type Camera struct {
    Position mgl32.Vec3
    Front    mgl32.Vec3
    Up       mgl32.Vec3
    Right    mgl32.Vec3
    WorldUp  mgl32.Vec3
    
    Yaw      float32
    Pitch    float32
}

func NewCamera(position, up mgl32.Vec3, yaw, pitch float32) *Camera {
    cam := &Camera{
        Position: position,
        WorldUp:  up,
        Yaw:      yaw,
        Pitch:    pitch,
        Front:    mgl32.Vec3{0, 0, -1},
    }
    cam.updateCameraVectors()
    return cam
}

func (c *Camera) GetViewMatrix() mgl32.Mat4 {
    return mgl32.LookAtV(c.Position, c.Position.Add(c.Front), c.Up)
}

func (c *Camera) updateCameraVectors() {
    // Calculate the new Front vector
    front := mgl32.Vec3{
        float32(math.Cos(float64(mgl32.DegToRad(c.Yaw))) * math.Cos(float64(mgl32.DegToRad(c.Pitch)))),
        float32(math.Sin(float64(mgl32.DegToRad(c.Pitch)))),
        float32(math.Sin(float64(mgl32.DegToRad(c.Yaw))) * math.Cos(float64(mgl32.DegToRad(c.Pitch)))),
    }
    c.Front = front.Normalize()
    c.Right = c.Front.Cross(c.WorldUp).Normalize()
    c.Up = c.Right.Cross(c.Front).Normalize()
}


func (c *Camera) ProcessKeyboard(direction string, deltaTime float32) {
    velocity := 2.5 * deltaTime
    
    switch direction {
    case "FORWARD":
        c.Position = c.Position.Add(c.Front.Mul(velocity))
    case "BACKWARD":
        c.Position = c.Position.Sub(c.Front.Mul(velocity))
    case "LEFT":
        c.Position = c.Position.Sub(c.Right.Mul(velocity))
    case "RIGHT":
        c.Position = c.Position.Add(c.Right.Mul(velocity))
    }
}

func (c *Camera) ProcessMouseMovement(xoffset, yoffset float32) {
    sensitivity := float32(0.1)
    xoffset *= sensitivity
    yoffset *= sensitivity
    
    c.Yaw += xoffset
    c.Pitch += yoffset
    
    // Constrain pitch
    if c.Pitch > 89.0 {
        c.Pitch = 89.0
    }
    if c.Pitch < -89.0 {
        c.Pitch = -89.0
    }
    
    c.updateCameraVectors()
}