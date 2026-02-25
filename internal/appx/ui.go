package appx

import (
	"github.com/tmazitov/42_scop/internal/geom"
	"github.com/tmazitov/42_scop/internal/clr"
	"github.com/tmazitov/42_scop/internal/ui"
	"github.com/go-gl/gl/v2.1/gl"
)

func (a *App) SetupButtons() {
	a.ui.AddButton(ui.NewButton().	
		SetPos(&geom.Pos{X: 10, Y: 10, Z: 1}).
		SetSize(40, 40).
		SetColor(clr.NewColor(0, 0, 255)).
		SetOnClick(func (xpos, ypos float32) error {
			
			a.state.IsVertexOnly = !a.state.IsVertexOnly

			if a.state.IsVertexOnly {
				gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
			} else {
				gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
			}

			return nil
		}))
}