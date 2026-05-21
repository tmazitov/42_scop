package controller

import (
	"github.com/tmazitov/42_scop/internal/appx/camera"
	"github.com/tmazitov/42_scop/internal/appx/window"
	"github.com/tmazitov/42_scop/internal/ui"
	"github.com/tmazitov/42_scop/internal/rende"
)

type App interface {
	Camera() *camera.Camera
	UI() *ui.UI
	Window() *window.Window
	ScreenSize() rende.ScreenSize
	TranslateSelectedObject(dx, dy, dz float32)
	RotateSelectedObject(angleX, angleY, angleZ float32)
}