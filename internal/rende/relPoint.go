package rende

type RelativePoint struct {
	X float32
	Y float32
	Z float32
}

func (p *RelativePoint) ConvertToVertex(screen ScreenSize) *Vertex {
	relativeX := convertX(p.X, screen.Width / 2)
	relativeY := convertY(p.Y, screen.Height / 2)
	return NewVertex([3]float32{
		relativeX,
		relativeY,
		0,
	})	
}

func convertY(value, center float32) float32{

	if value == center {
		return 0
	}

	if value < center {
		return value / center
	}
	
	return -1 * ((value / center) - 1)
}

func convertX(value, center float32) float32{

	if value == center {
		return 0
	}

	if value < center {
		return -1 * value / center
	}
	
	return ((value / center) - 1)
}