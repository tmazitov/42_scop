package ui

import (
	// "github.com/tmazitov/42_scop/internal/rende"
)

type UI struct {
	buttons []*Button
	texts	[]*Text
}

func NewUI() *UI {
	return &UI{
		buttons: nil,
		texts: nil,
	}
}

func (ui *UI) AddButton(button *Button) {
	ui.buttons = append(ui.buttons, button)
}

func (ui *UI) AddStaticText(text *Text) {
	ui.texts = append(ui.texts, text)
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
	for _, text := range ui.texts {
		text.Draw()
	}
}

func (ui *UI) Cleanup(){
	for _, text := range ui.texts {
		text.Cleanup()
	}
}

