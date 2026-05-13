package appx

import (
	"github.com/tmazitov/42_scop/internal/geom"
	"github.com/tmazitov/42_scop/internal/clr"
	"github.com/tmazitov/42_scop/internal/ui"
)

func (a *App) SetupButtons() {

	enabledColor := clr.NewColor(245, 193, 24)
	disabledColor := clr.NewColor(54, 50, 40)

	buttons := []*ui.Button{

		// Toggle Vertex/Fill mode button
		ui.NewButton().	
		SetPos(&geom.Pos{X: 10, Y: 10, Z: 1}).
		SetSize(40, 120).
		SetColor(enabledColor).
		SetText("Vertex Mode").
		SetOnClick(func (instance *ui.Button, xpos, ypos float32) error {
			
			a.state.IsVertexOnly = !a.state.IsVertexOnly

			if a.state.IsVertexOnly {
				instance.SetColor(disabledColor).SetText("Fill Mode")
			} else {
				instance.SetColor(enabledColor).SetText("Vertex Mode")
			}

			return nil
		}),

		// Toggle Rotation button

		// Toggle Vertex/Fill mode button
		ui.NewButton().	
		SetPos(&geom.Pos{X: 10, Y: 60, Z: 1}).
		SetSize(40, 120).
		SetColor(enabledColor).
		SetText("Enable Rotation").
		SetOnClick(func (instance *ui.Button, xpos, ypos float32) error {
			
			a.state.IsRotationEnabled = !a.state.IsRotationEnabled

			if a.state.IsRotationEnabled {
				instance.SetColor(disabledColor).SetText("Disable Rotation")
			} else {
				instance.SetColor(enabledColor).SetText("Enable Rotation")
			}

			return nil
		}),
	}

	for _, btn := range buttons {
		a.ui.AddButton(btn)
	}
}