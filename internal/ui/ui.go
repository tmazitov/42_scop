package ui

import (
	"github.com/tmazitov/42_scop/internal/rende"
)

type UI struct {
	buttons []*Button
	screenSize rende.ScreenSize
}

func NewUI(screenSize rende.ScreenSize) *UI {
	return &UI{
		buttons: nil,
		screenSize: screenSize,
	}
}

func (ui *UI) AddButton(button *Button) {
	ui.buttons = append(ui.buttons, button)
}

func (ui *UI) IsPressed(xpos, ypos float32) ElementHandleFunc {
	for _, button := range ui.buttons {
		if button.IsPressed(xpos, ypos) {
			return button.OnClickHandler()
		}
	}
	return nil
}

func (ui *UI) Draw() {
	for _, button := range ui.buttons{
		button.Draw()
	}
}

