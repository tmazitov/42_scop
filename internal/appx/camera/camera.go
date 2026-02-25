package camera

import (
    "github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	Position mgl32.Vec3
	Front    mgl32.Vec3		// Norm vector : Determines what the camera sees
	Up       mgl32.Vec3		// Norm vector : Defines which way is "up" for camera orientation
	Right    mgl32.Vec3		// Norm vector : Used for left/right strafing movement
	WorldUp  mgl32.Vec3		// Norm vector : Reference for calculating camera's local coordinate system

	Yaw   float32			// Horizontal camera rotation in degrees (-90 - forward, 0 -right, 180 - left)
	Pitch float32			// Vertical camera rotation in degrees ()
	Zoom  float32

	MovementSpeed    float32
	MouseSensitivity float32
}

func NewCamera(position mgl32.Vec3) *Camera {
	cam := &Camera{
		Position: position,
		WorldUp:  mgl32.Vec3{0, 1, 0},
		Front:    mgl32.Vec3{0, 0, -1},
		Zoom:	  1,
	}
	cam.updateCameraVectors()
	return cam
}

